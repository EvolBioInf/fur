package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/esa"
	"github.com/evolbioinf/fur/util"
	"github.com/ivantsers/fasta"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
)

func readDir(dir string) map[string]bool {
	dirEntries, err := os.ReadDir(dir)
	util.Check(err)
	names := make(map[string]bool)
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			p := dir + "/" + dirEntry.Name()
			fmt.Fprintf(os.Stderr,
				"skipping subdirectory %s\n", p)
			continue
		}
		ext := filepath.Ext(dirEntry.Name())
		if ext != ".fasta" && ext != ".fna" && ext != ".ffn" &&
			ext != ".frn" && ext != ".fa" {
			m := "%s doesn't have the extension of " +
				"a nucleotide FASTA file; skipping it\n"
			p := dir + "/" + dirEntry.Name()
			fmt.Fprintf(os.Stderr, m, p)
			continue
		}
		names[dirEntry.Name()] = true
	}
	return names
}
func matchSeq(d []byte, e *esa.Esa, ml []int, rev bool) {
	for i := 0; i < len(d); {
		match := e.MatchPref(d[i:])
		if match.L == 0 {
			match.L = 1
		}
		p := i
		if rev {
			p = len(d) - i - match.L
		}
		if ml[p] < match.L {
			ml[p] = match.L
		}
		i += match.L + 1
	}
}
func main() {
	util.PrepareErrorMessages("makeFurDb")
	optV := flag.Bool("v", false, "version")
	optT := flag.String("t", "", "target directory")
	optN := flag.String("n", "", "neighbor directory")
	optD := flag.String("d", "", "database directory")
	optR := flag.String("r", "", "target representative "+
		"(default shortest)")
	optO := flag.Bool("o", false, "overwrite existing database")
	optTT := flag.Int("T", 0, "threads (default all processors)")
	u := "makeFurDb [option]... -t <targetDir> " +
		"-n <neighborDir> -d <db>"
	p := "Construct fur database."
	e := "makeFurDb -t targets/ -n neighbors/ -d fur.db"
	clio.Usage(u, p, e)
	flag.Parse()
	if *optV {
		util.PrintInfo("makeFurDb")
	}
	if *optT == "" {
		m := "please provide a directory " +
			"of target sequences"
		fmt.Fprintf(os.Stderr, "%s\n", m)
		os.Exit(1)
	}
	if *optN == "" {
		m := "please provide a directory " +
			"of neighbor sequences"
		fmt.Fprintf(os.Stderr, "%s\n", m)
		os.Exit(1)
	}
	if *optD == "" {
		m := "please provide a database name"
		fmt.Fprintf(os.Stderr, "%s\n", m)
		os.Exit(1)
	} else {
		_, err := os.Stat(*optD)
		if err == nil {
			if *optO {
				err := os.RemoveAll(*optD)
				util.Check(err)
			} else {
				m := fmt.Sprintf("database %s already exists", *optD)
				fmt.Fprintf(os.Stderr, "%s\n", m)
				os.Exit(1)
			}
		}
		err = os.Mkdir(*optD, 0750)
		util.Check(err)
	}

	if *optTT < 0 {
		log.Fatalf("Can't set %d threads.", *optTT)
	}
	if *optTT == 0 {
		(*optTT) = runtime.NumCPU()
	}
	targets := readDir(*optT)
	if len(targets) == 0 {
		fmt.Fprintf(os.Stderr, "%s is empty\n", *optT)
		os.Exit(1)
	}
	neighbors := readDir(*optN)
	if len(neighbors) == 0 {
		fmt.Fprintf(os.Stderr, "%s is empty\n", *optN)
		os.Exit(1)
	}
	var targetNames, neighborNames []string
	for target := range targets {
		targetNames = append(targetNames, target)
	}
	sort.Strings(targetNames)
	for neighbor := range neighbors {
		neighborNames = append(neighborNames, neighbor)
	}
	sort.Strings(neighborNames)
	for _, target := range targetNames {
		if neighbors[target] {
			m := "found %s/%s and %s/%s; please " +
				"make sure the targets and " +
				"neighbors don't overlap"
			fmt.Fprintf(os.Stderr, m, *optT,
				target, *optN, target)
			os.Exit(1)
		}
	}
	vf := *optD + "/v.txt"
	f, err := os.Create(vf)
	util.Check(err)
	fmt.Fprintf(f, "%s\n", util.Version())
	f.Close()
	if *optR == "" {
		minTar := ""
		minLen := math.MaxInt
		for target, _ := range targets {
			l := 0
			p := *optT + "/" + target
			f, err := os.Open(p)
			util.Check(err)
			sc := fasta.NewScanner(f)
			for sc.ScanSequence() {
				l += len(sc.Sequence().Data())
			}
			f.Close()
			if l < minLen {
				minTar = target
				minLen = l
			}
		}
		(*optR) = minTar
	}
	fmt.Fprintf(os.Stderr, "using %s as target representative\n",
		(*optR))
	var repSeqs, revRepSeqs []*fasta.Sequence
	p = *optT + "/" + *optR
	f, err = os.Open(p)
	util.Check(err)
	defer f.Close()
	sc := fasta.NewScanner(f)
	for sc.ScanSequence() {
		seq := sc.Sequence()
		h := strings.Fields(seq.Header())[0]
		seq = fasta.NewSequence(h, seq.Data())
		repSeqs = append(repSeqs, seq)
		seq = fasta.NewSequence(seq.Header(), seq.Data())
		seq.ReverseComplement()
		revRepSeqs = append(revRepSeqs, seq)
	}
	f, err = os.Create(*optD + "/r.fasta")
	util.Check(err)
	defer f.Close()
	for _, repSeq := range repSeqs {
		fmt.Fprintf(f, "%s\n", repSeq)
	}
	p = *optD + "/t/"
	err = os.Mkdir(p, 0750)
	util.Check(err)
	for target, _ := range targets {
		if target == *optR {
			continue
		}
		source := *optT + "/" + target
		dest := *optD + "/t/" + target
		sd, err := os.ReadFile(source)
		util.Check(err)
		err = os.WriteFile(dest, sd, 0666)
	}
	f, err = os.Open(*optD + "/r.fasta")
	util.Check(err)
	var targetSeqs, revTargetSeqs []*fasta.Sequence
	sc = fasta.NewScanner(f)
	for sc.ScanSequence() {
		s := sc.Sequence()
		targetSeqs = append(targetSeqs, s)
		s = fasta.NewSequence(s.Header(), s.Data())
		s.ReverseComplement()
		revTargetSeqs = append(revTargetSeqs, s)
	}
	f.Close()
	for i, targetSeq := range targetSeqs {
		h := targetSeq.Header()
		d := bytes.ToUpper(targetSeq.Data())
		targetSeqs[i] = fasta.NewSequence(h, d)
		h = revTargetSeqs[i].Header()
		d = bytes.ToUpper(revTargetSeqs[i].Data())
		revTargetSeqs[i] = fasta.NewSequence(h, d)
	}
	mnl := -1
	for _, neighbor := range neighborNames {
		p := *optN + "/" + neighbor
		f, err := os.Open(p)
		util.Check(err)
		sc := fasta.NewScanner(f)
		for sc.ScanSequence() {
			l := len(sc.Sequence().Data())
			if l > mnl {
				mnl = l
			}
		}
		f.Close()
	}
	neighborNameSets := make([][]string, 0)
	n := len(neighborNames)
	length := int(math.Ceil(float64(n) / float64(*optTT)))
	start := 0
	end := length
	for start < n {
		neighborNameSets = append(neighborNameSets,
			neighborNames[start:end])
		start = end
		end += length
		if end > n {
			end = n
		}
	}
	lengthSets := make(chan [][]int)
	var wg sync.WaitGroup
	for _, neighborNames := range neighborNameSets {
		wg.Add(1)
		go func(neighborNames []string) {
			defer wg.Done()
			var matchLengths [][]int
			for _, targetSeq := range targetSeqs {
				n := len(targetSeq.Data())
				lengths := make([]int, n)
				matchLengths = append(matchLengths, lengths)
			}
			li := 0
			for _, neighbor := range neighborNames {
				p := *optN + "/" + neighbor
				f, err := os.Open(p)
				util.Check(err)
				sc := fasta.NewScanner(f)
				for sc.ScanSequence() {
					s := sc.Sequence()
					d := s.Data()
					h := []byte(s.Header())
					for len(d) < mnl && sc.ScanSequence() {
						s = sc.Sequence()
						h = append(h, '|')
						h = append(h, []byte(s.Header())...)
						d = append(d, s.Data()...)
						li++
					}
					d = bytes.ToUpper(d)
					e := esa.MakeEsa(d)
					for i, targetSeq := range targetSeqs {
						d := targetSeq.Data()
						rev := false
						matchSeq(d, e, matchLengths[i], rev)
						d = revTargetSeqs[i].Data()
						rev = true
						matchSeq(d, e, matchLengths[i], rev)
					}
				}
				f.Close()
			}
			lengthSets <- matchLengths
		}(neighborNames)
	}
	go func() {
		wg.Wait()
		close(lengthSets)
	}()
	matchLengths := make([][]int, 0)
	for _, ts := range targetSeqs {
		ml := make([]int, len(ts.Data()))
		matchLengths = append(matchLengths, ml)
	}
	for lengthSet := range lengthSets {
		for i, lengths := range lengthSet {
			for j, length := range lengths {
				if matchLengths[i][j] < length {
					matchLengths[i][j] = length
				}
			}
		}
	}
	for _, ml := range matchLengths {
		l := 0
		for i := 0; i < len(ml); i++ {
			if ml[i] > l {
				l = ml[i]
			}
			ml[i] = l
			l--
		}
	}
	for _, ml := range matchLengths {
		for i := 0; i < len(ml); {
			m := ml[i]
			for j := 0; j < m-1; j++ {
				ml[i+j] = 0
			}
			ml[i+m-1] = 1
			i += m + 1
		}
	}
	f, err = os.Create(*optD + "/e.fasta")
	util.Check(err)
	wr := bufio.NewWriter(f)
	for i, seq := range targetSeqs {
		ml := matchLengths[i]
		fmt.Fprintf(wr, ">%s\n", seq.Header())
		for j := 0; j < len(ml); j++ {
			c := '0'
			if ml[j] == 1 {
				c = '1'
			}
			fmt.Fprintf(wr, "%c", c)
			if (j+1)%fasta.DefaultLineLength == 0 {
				fmt.Fprintf(wr, "\n")
			}
		}
		if len(ml)%fasta.DefaultLineLength != 0 {
			fmt.Fprintf(wr, "\n")
		}
	}
	err = wr.Flush()
	util.Check(err)
	f.Close()
	fmt.Fprintf(os.Stderr, "making Blast database\n")
	mask := *optD + "/mask.asnb"
	cmd := exec.Command("convert2blastmask",
		"-masking_algorithm", "repeat",
		"-masking_options", "default",
		"-outfmt", "maskinfo_asn1_bin",
		"-out", mask)
	stdin, err := cmd.StdinPipe()
	util.Check(err)
	go func() {
		defer stdin.Close()
		for neighbor, _ := range neighbors {
			p := *optN + "/" + neighbor
			d, err := os.ReadFile(p)
			util.Check(err)
			fmt.Fprint(stdin, string(d))
		}
	}()
	_, err = cmd.Output()
	util.Check(err)
	cmd = exec.Command("makeblastdb",
		"-dbtype", "nucl",
		"-out", *optD+"/n",
		"-title", "n",
		"-mask_data", mask)
	stdin, err = cmd.StdinPipe()
	util.Check(err)
	go func() {
		defer stdin.Close()
		for neighbor, _ := range neighbors {
			p := *optN + "/" + neighbor
			d, err := os.ReadFile(p)
			util.Check(err)
			fmt.Fprint(stdin, string(d))
		}
	}()
	_, err = cmd.Output()
	util.Check(err)
	w, err := os.Create(*optD + "/n.txt")
	util.Check(err)
	defer w.Close()
	cmd = exec.Command("blastdbcmd", "-db",
		(*optD)+"/n",
		"-entry", "all")
	out, err := cmd.Output()
	util.Check(err)
	r := bytes.NewReader(out)
	sc = fasta.NewScanner(r)
	var l, g int
	nuc := "ACGTacgt"
	gc := "GCgc"
	for sc.ScanSequence() {
		d := sc.Sequence().Data()
		for i, _ := range d {
			if bytes.ContainsAny(d[i:i+1], nuc) {
				l++
			} else {
				continue
			}
			if bytes.ContainsAny(d[i:i+1], gc) {
				g++
			}
		}
	}
	gcc := float64(g) / float64(l)
	fmt.Fprintf(w, "length: %d\nGC-content: %f\n", l, gcc)
}
