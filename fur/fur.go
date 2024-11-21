package main

import (
          "github.com/evolbioinf/fur/util"
          "github.com/evolbioinf/clio"
          "flag"
          "runtime"
          "fmt"
          "os"
          "strings"
          "strconv"
          "log"
          "github.com/evolbioinf/fasta"
          "text/tabwriter"
          "github.com/evolbioinf/sus"
          "math"
          "path/filepath"
          "os/exec"
          "bytes"
)
type interval struct {
          s, e int
}
func countNucl(sequences []*fasta.Sequence) (l, n int) {
          for _, sequence := range sequences {
                  l += len(sequence.Data())
                  for _, c := range sequence.Data() {
                          if c == 'N' { n++ }
                  }
          }
          return l, n
}
func main() {
          util.PrepareErrorMessages("fur")
          u := "fur -d <db> [option]..."
          p := "Find unique regions."
          e := "fur -d fur.db"
          clio.Usage(u, p, e)
          optV := flag.Bool("v", false, "version")
          optD := flag.String("d", "", "database")
          optW := flag.Int("w", 80, "window length")
          m := "quantile of match length distribution"
          optQ := flag.Float64("q", 0.1, m)
          m = "print unique regions after sliding window analysis " +
                    "and exit"
          optU := flag.Bool("u", false, m)
          m = "print unique regions after checking for presence " +
                    "in templates and exit"
          optX := flag.Bool("x", false, "exact matches only")
          optUU := flag.Bool("U", false, m)
          optE := flag.Float64("e", 1e-5, "E-value for Blast")
          ncpu := runtime.NumCPU()
          optT := flag.Int("t", ncpu, "Number of threads " +
                    "for Phylonium and Blast")
          optM := flag.Bool("m", false, "megablast mode " +
                    "(default blastn)")
          optMM := flag.Bool("M", false,
                    "activate masking (recommended for mammalian genomes)")
          optN := flag.Int("n", 100, "number of nucleotides in region")
          flag.Parse()
          if *optV {
                    util.PrintInfo("fur")
          }
          if *optD == "" {
                    fmt.Fprintf(os.Stderr, "please supply database\n")
                    os.Exit(1)
          } else {
                    _, err := os.Stat(*optD)
                    util.Check(err)
          }
          db, err := os.ReadFile(*optD + "/v.txt")
          util.Check(err)
          ds := strings.Split(string(db[1:]), "-")[0]
          ds = strings.TrimRight(ds, "\n")
          dv, err := strconv.ParseFloat(ds, 64)
          if err != nil {
                    m := "couldn't read the datatase version from %q"
                    log.Fatalf(m, ds)
          }
          vs := util.Version()
          ps := ""
          if len(vs) > 0 {
                    ps = strings.Split(vs[1:], "-")[0]
          }
          pv, err := strconv.ParseFloat(ps, 64)
          util.Check(err)
          m = "fur v%s is incompatible with database v%s\n"
          if pv < dv {
                    fmt.Fprintf(os.Stderr, m, ps, ds)
                    os.Exit(1)
          }
          if *optT > ncpu {
                    m := "Warning [fur]: Number of threads was reduced " +
                            "to %d to match the number of available " +
                            "CPUs.\n"
                    fmt.Fprintf(os.Stderr, m, ncpu)
                    (*optT) = ncpu
          }
          regions := make([]*fasta.Sequence, 0)
          rw := tabwriter.NewWriter(os.Stderr, 0, 0, 2, ' ',
                    tabwriter.AlignRight)
          var ends []*fasta.Sequence
          f, err := os.Open(*optD + "/e.fasta")
          util.Check(err)
          sc := fasta.NewScanner(f)
          for sc.ScanSequence() {
                    ends = append(ends, sc.Sequence())
          }
          f.Close()
          d, err := os.ReadFile(*optD + "/n.txt")
          util.Check(err)
          fields := strings.Fields(string(d))
          l, err := strconv.Atoi(fields[1])
          util.Check(err)
          g, err := strconv.ParseFloat(fields[3], 64)
          util.Check(err)
          q := sus.Quantile(l, g, *optQ) - 1
          t := int(math.Round(float64(*optW) / float64(q)))
          intervals := make([][]*interval, len(ends))
          for i, end := range ends {
                    d := end.Data()
                    if len(d) >= *optW {
                            nm := 1
                            l := 0
                            r := 0
                            for r < *optW {
                                      if d[r] == '1' {
                                              nm++
                                      }
                                      r++
                            }
                            open := false
                            var iv *interval
                            for r < len(d) {
                                      if nm >= t {
                                                if open {
                                                          iv.e = r
                                                } else {
                                                          iv = new(interval)
                                                          iv.s = l
                                                          iv.e = r
                                                          open = true
                                                }
                                      } else if open && iv.e < l {
                                                open = false
                                                intervals[i] = append(intervals[i], iv)
                                      }
                                      if d[l] == '1' {
                                                nm--
                                      }
                                      l++
                                      if d[r] == '1' {
                                                nm++
                                      }
                                      r++
                            }
                            if open {
                                      intervals[i] = append(intervals[i], iv)
                            }
                    }
          }
          f, err = os.Open(*optD + "/r.fasta")
          util.Check(err)
          seqAcc := make(map[string]bool)
          sc = fasta.NewScanner(f)
          var r []*fasta.Sequence
          for sc.ScanSequence() {
                    s := sc.Sequence()
                    acc := strings.Fields(s.Header())[0]
                    if seqAcc[acc] {
                              log.Fatalf("%q is not a unique accession", acc)
                    }
                    seqAcc[acc] = true
                    r = append(r, s)
          }
          for i, interval := range intervals {
                    d := r[i].Data()
                    for _, iv := range interval {
                            l := iv.e - iv.s + 1
                            if l >= *optN {
                                    arr := strings.Fields(r[i].Header())
                                    h := fmt.Sprintf("%s_(%d..%d)",
                                              arr[0], iv.s+1, iv.e+1)
                                    region := fasta.NewSequence(h, d[iv.s:iv.e+1])
                                    regions = append(regions, region)
                            }
                    }
          }
          fmt.Fprintf(rw, "%s\t%s\t%s\t%s\t\n", "Step         ",
                    "Sequences", "Length", "Ns")
          fmt.Fprintf(rw, "%s\t%s\t%s\t%s\t\n", "-------------",
                    "---------", "------", "--")
          rf := "%s\t%d\t%d\t%d\t\n"
          ns := len(regions)
          le, nn := countNucl(regions)
          fmt.Fprintf(rw, rf, "Subtraction_1", ns, le, nn)
          if len(regions) == 0  || *optU {
                    rw.Flush()
          }
          if *optU {
                    for _, region := range regions {
                            fmt.Printf("%s\n", region)
                    }
                    os.Exit(0)
          }
          if len(regions) > 0 {
                    numTargets := 0
                    dirEntries, err := os.ReadDir(*optD + "/t")
                    numTargets = len(dirEntries) + 1
                    if numTargets > 1 {
                              f, err = os.CreateTemp(*optD, "*.fasta")
                              util.Check(err)
                              for _, region := range regions {
                                        fmt.Fprintf(f, "%s\n", region)
                              }
                              f.Close()
                              cmd := exec.Command("phylonium")
                              rf := f.Name()
                              pf := ""
                              f, err = os.CreateTemp(*optD, "*.fasta")
                              util.Check(err)
                              pf = f.Name()
                              f.Close()
                              _, err := os.Stat(pf)
                              if err == nil {
                                        err = os.Remove(pf)
                                        util.Check(err)
                              }
                              ts := strconv.Itoa(*optT)
                              args := []string{"phylonium", "-t", ts, "-p", pf, "-r", rf}
                              tf, err := filepath.Glob(*optD + "/t/*")
                              args = append(args, tf...)
                              cmd.Args = args
                              out, err := cmd.CombinedOutput()
                              if err != nil {
                                        i := bytes.Index(out, []byte("homology"))
                                        j := bytes.Index(out, []byte("nan"))
                                        if i < 0  && j < 0 {
                                                fmt.Fprintf(os.Stderr,"%s\n", out)
                                                os.Exit(1)
                                        }
                              }
                              regions = regions[:0]
                              f, err = os.Open(pf)
                              util.Check(err)
                              sc = fasta.NewScanner(f)
                              for sc.ScanSequence() {
                                        s := sc.Sequence()
                                        regions = append(regions, s)
                              }
                              f.Close()
                              i := 0
                              for _, region := range regions {
                                        if len(region.Data()) >= *optN {
                                                regions[i] = region
                                                i++
                                        }
                              }
                              regions = regions[:i]
                              i = 0
                              for _, region := range regions {
                                        arr := strings.Fields(region.Header())
                                        n, err := strconv.Atoi(arr[2])
                                        util.Check(err)
                                        pos := make([]int, 0)
                                        for j := 0; j < n; j++ {
                                                  x, err := strconv.Atoi(arr[j+3])
                                                  util.Check(err)
                                                  pos = append(pos, x-1)
                                        }
                                        if *optX && len(pos) > 0 { continue }
                                        for _, p := range pos {
                                                  region.Data()[p] = 'N'
                                        }
                                        regions[i] = region
                                        i++
                              }
                              regions = regions[:i]
                              for _, region := range regions {
                                        arr := strings.Split(region.Header(), "(")
                                        arr = strings.Split(arr[1], ")")
                                        mutations := arr[1]
                                        arr = strings.Split(arr[0], "..")
                                        sp, err := strconv.Atoi(arr[0])
                                        util.Check(err)
                                        ep, err := strconv.Atoi(arr[1])
                                        util.Check(err)
                                        sp -= 1
                                        ep -= 2
                                        var sr, er, ii, iii int
                                        c := 0
                                        found := false
                                        intervals:
                                        for ii < len(intervals) {
                                                  ivals := intervals[ii]
                                                  for iii < len(ivals) {
                                                          ival := ivals[iii]
                                                          iii++
                                                          l := ival.e - ival.s + 1
                                                          if l < *optN { continue }
                                                          if sp <= c+l-1 && ep >= c {
                                                                    sr = ival.s + sp - c
                                                                    er = ival.s + ep - c
                                                                    found = true
                                                                    c += l+1
                                                                    break intervals
                                                          }
                                                          c += l+1
                                                  }
                                                  iii = 0
                                                  ii++
                                        }
                                        if found {
                                                  h := fmt.Sprintf("%s_(%d..%d) %s", r[ii].Header(),
                                                          sr+1, er+1, mutations)
                                                  region.SetHeader(h)
                                        } else {
                                                  log.Fatalf("Coudn't find region %s\n",
                                                          region.Header())
                                        }
                              }
                              for _, region := range regions {
                                        for i, c := range region.Data() {
                                                if c == '!' {
                                                        region.Data()[i] = 'N'
                                                }
                                        }
                              }
                              err = os.Remove(rf)
                              util.Check(err)
                              err = os.Remove(pf)
                              util.Check(err)
                    }
                    ns = len(regions)
                    le, nn = countNucl(regions)
                    fmt.Fprintf(rw, rf, "Intersection ", ns, le, nn)
                    if len(regions) == 0 || *optUU {
                              rw.Flush()
                    }
                    if *optUU {
                              for _, region := range regions {
                                      fmt.Printf("%s\n", region)
                              }
                              os.Exit(0)
                    }
          }
          if len(regions) > 0 {
                    cmds := make([]*exec.Cmd, 0)
                    da := *optD + "/n"
                    th := *optT
                    ev := *optE
                    ta := "megablast"
                    ma := ""
                    if *optMM {
                              cmd := exec.Command("blastdbcmd", "-info", "-db", *optD + "/n")
                              info, err := cmd.CombinedOutput()
                              util.Check(err)
                              lines := strings.Split(string(info), "\n")
                              for i, line := range lines {
                                        fields := strings.Fields(line)
                                        if len(fields) > 0 && fields[0] == "Algorithm" {
                                                ma = strings.Fields(lines[i+1])[0]
                                        }
                              }
                              if ma == "" {
                                        m := "#Warning [fur]: No masking information " +
                                                "in Blast database; running Subtraction_2 " +
                                                "without masking.\n"
                                        fmt.Fprintf(os.Stderr, m)
                              }
                    }
                    of := "6 qaccver qstart qend"
                    tm := "blastn -db %s -num_threads %d "
                    tm += "-evalue %g -task %s "
                    if *optMM  && ma != "" {
                              tm += "-db_soft_mask %s "
                    }
                    tm += "-outfmt "
                    as := fmt.Sprintf(tm, da, th, ev, ta)
                    if *optMM && ma != "" {
                              as = fmt.Sprintf(tm, da, th, ev, ta, ma)
                    }
                    args := strings.Fields(as)
                    args = append(args, of)
                    cmd := exec.Command("blastn")
                    cmd.Args = args
                    cmds = append(cmds, cmd)
                    if !*optM {
                              ta = "blastn"
                              as := fmt.Sprintf(tm, da, th, ev, ta)
                              if *optMM && ma != "" {
                                        as = fmt.Sprintf(tm, da, th, ev, ta, ma)
                              }
                              args = strings.Fields(as)
                              args = append(args, of)
                              cmd = exec.Command("blastn")
                              cmd.Args = args
                              cmds = append(cmds, cmd)
                    }
                    for _, cmd := range cmds {
                              stdin, err := cmd.StdinPipe()
                              util.Check(err)
                              go func() {
                                        defer stdin.Close()
                                        for _, region := range regions {
                                                fmt.Fprintf(stdin, "%s\n", region)
                                        }
                              }()
                              b, err := cmd.CombinedOutput()
                              if err != nil {
                                        log.Fatalf("%s\n", string(b))
                              }
                              hits := bytes.Split(b, []byte("\n"))
                              hits = hits[:len(hits)-1]
                              regMap := make(map[string]int)
                              le = 0
                              for i, region := range regions {
                                        le += len(region.Data())
                                        acc := strings.Fields(region.Header())[0]
                                        regMap[acc] = i
                              }
                              for _, hit := range hits {
                                        arr := strings.Fields(string(hit))
                                        if len(arr) != 3 {
                                                  log.Fatalf("Failed Blast: %s\n", string(hit))
                                        }
                                        qacc := arr[0]
                                        qstart, err := strconv.Atoi(arr[1])
                                        util.Check(err)
                                        qend, err := strconv.Atoi(arr[2])
                                        util.Check(err)
                                        i := regMap[qacc]
                                        r := regions[i].Data()
                                        qstart--
                                        qend--
                                        offset := 15
                                        if qstart < offset {
                                                  qstart = 0
                                        }
                                        if qend > len(r) - offset - 1 {
                                                  qend = len(r) - 1
                                        }
                                        for i := qstart; i <= qend; i++ {
                                                  r[i] = 'N'
                                        }
                              }
                              for i, region := range regions {
                                        h := region.Header()
                                        r := bytes.TrimLeft(region.Data(), "N")
                                        dl := len(region.Data()) - len(r)
                                        r = bytes.TrimRight(r, "N")
                                        dr := len(region.Data()) - len(r) - dl
                                        if dl > 0 || dr > 0 {
                                                arr := strings.Split(h, "_(")
                                                prefix := arr[0]
                                                arr = strings.Split(arr[1], ")")
                                                muts := strings.Fields(arr[1])
                                                mutations := make([]int, 0)
                                                for _, m := range muts {
                                                          i, err := strconv.Atoi(m)
                                                          util.Check(err)
                                                          mutations = append(mutations, i)
                                                }
                                                arr = strings.Split(arr[0], "..")
                                                s, err := strconv.Atoi(arr[0])
                                                util.Check(err)
                                                e, err := strconv.Atoi(arr[1])
                                                util.Check(err)
                                                s += dl
                                                e -= dr
                                                nm := make([]int, 0)
                                                l := e - s
                                                for i := 1; i < len(mutations); i++ {
                                                          x := mutations[i] - dl
                                                          if x > 0 && x <= l {
                                                                  nm = append(nm, x)
                                                          }
                                                }
                                                n := len(nm)
                                                h = fmt.Sprintf("%s_(%d..%d) %4d",
                                                          prefix, s, e, n)
                                                for _, m := range nm {
                                                          h = fmt.Sprintf("%s %d", h, m)
                                                }
                                        }
                                        s := fasta.NewSequence(h, r)
                                        regions[i] = s
                              }
                              i := 0
                              sa := make([]*fasta.Sequence, 1)
                              for _, region := range regions {
                                        sa[0] = region
                                        l, n := countNucl(sa)
                                        if l-n >= *optN {
                                                regions[i] = region
                                                i++
                                        }
                              }
                              regions = regions[:i]
                    }
                    ns = len(regions)
                    le, nn = countNucl(regions)
                    fmt.Fprintf(rw, rf, "Subtraction_2", ns, le, nn)
                    rw.Flush()
                    for _, region := range regions {
                              fmt.Printf("%s\n", region)
                    }
          }
}
