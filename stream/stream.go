package main

import (
          "github.com/evolbioinf/fur/util"
          "github.com/evolbioinf/clio"
          "flag"
          "log"
          "time"
          "math/rand"
          "io"
          "github.com/ivantsers/fasta"
          "fmt"
          "github.com/evolbioinf/esa"
)
func scan(r io.Reader, args ...interface{}) {
          ra := args[0].(*rand.Rand)
          optI := args[1].(*bool)
          optM := args[2].(*float64)
          optS := args[3].(*float64)
          sc := fasta.NewScanner(r)
          for sc.ScanSequence() {
                  s1 := sc.Sequence()
                  s1f := s1.Data()
                  rev := fasta.NewSequence(s1.Header(), s1f)
                  rev.ReverseComplement()
                  s1fr := append(s1f, rev.Data()...)
                  s2f := make([]byte, len(s1f))
                  dic := []byte("ACGT")
                  for i, c := range s1f {
                            if ra.Float64() <= *optM {
                                    c = dic[ra.Intn(4)]
                            }
                            s2f[i] = c
                  }
                  rev = fasta.NewSequence(s1.Header(), s2f)
                  rev.ReverseComplement()
                  s2fr := append(s2f, rev.Data()...)
                  mf := 0
                  if *optI {
                            mf = compare(s1fr, s2f, *optS)
                  } else {
                            mf = compare(s1f, s2fr, *optS)
                  }
                  fmt.Printf("Match factors: %d\n", mf)
          }
}
func compare(s1, s2 []byte, skip float64) int {
          e := esa.MakeEsa(s1)
          mf := 0
          i := 0
          for i < len(s2) {
                    mf++
                    m := e.MatchPref(s2[i:])
                    s := int(float64(m.L) * skip)
                    if s == 0 {
                            s = 1
                    }
                    i += s
          }
          return mf
}
func main() {
          util.PrepareErrorMessages("stream")
          u := "stream [option]..."
          p := "Investigate streaming vs. indexing."
          e := "ranseq | stream"
          clio.Usage(u, p, e)
          optV := flag.Bool("v", false, "version")
          optI := flag.Bool("i", false, "indexing scenario " +
                    "(default streaming)")
          optS := flag.Float64("s", 0, "skipping fraction " +
                    "(default advance one base)")
          optM := flag.Float64("m", 0.01, "mutation rate")
          optSS := flag.Int64("S", 0, "seed for random number generator " +
                    "(default internal)")
          flag.Parse()
          if *optV {
                    util.PrintInfo("stream")
          }
          if *optS < 0 {
                    log.Fatal("plase use skipping fraction >= 0")
          }
          if *optM < 0 {
                    log.Fatal("please use mutation rate >= 0")
          }
          if *optSS == 0 {
                    (*optSS) = time.Now().UnixNano()
          }
          ra := rand.New(rand.NewSource(*optSS))
          files := flag.Args()
          clio.ParseFiles(files, scan, ra, optI, optM, optS)
}
