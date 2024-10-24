package main

import (
          "github.com/evolbioinf/fur/util"
          "flag"
          "github.com/evolbioinf/clio"
          "io"
          "github.com/ivantsers/fasta"
          "strconv"
          "fmt"
)
type interval struct {
          start, end int
}
func parse(r io.Reader, args ...interface{}) {
          maxRunLen := args[0].(int)
          minFragLen := args[1].(int)
          scanner := fasta.NewScanner(r)
          for scanner.ScanSequence() {
                  sequence := scanner.Sequence()
                  var runs []interval
                  var run interval
                  data := sequence.Data()
                  n := len(data)
                  for i := 0; i < n; i++ {
                            j := 0
                            for i+j < n && data[i+j] == 'N' { j++ }
                            if (i == 0 && j > 0) || i + j == n || j >= maxRunLen {
                                      run.start = i
                                      run.end = i + j - 1
                                      runs = append(runs, run)
                            }
                            i += j
                  }
                  var fragments []*fasta.Sequence
                  if len(runs) == 0 {
                            fragments = append(fragments, sequence)
                  } else {
                            prevEnd := -1
                            if runs[0].start == 0 && len(runs) > 1 {
                                      prevEnd = runs[0].end
                                      runs = runs[1:]
                            }
                            header := sequence.Header()
                            for _, run = range runs {
                                      seq := fasta.NewSequence(header,
                                              data[prevEnd+1:run.start])
                                      fragments = append(fragments, seq)
                                      prevEnd = run.end
                            }
                            if run.end < n {
                                      seq := fasta.NewSequence(header, data[run.end+1:])
                                      fragments = append(fragments, seq)
                            }
                  }
                  i := 0
                  for _, f := range fragments {
                            if len(f.Data()) >= minFragLen {
                                    fragments[i] = f
                                    i++
                            }
                  }
                  fragments = fragments[0:i]
                  if len(fragments) > 1 {
                            for i, f := range fragments {
                                    f.AppendToHeader(" - F" + strconv.Itoa(i+1))
                            }
                  }
                  for _, f := range fragments {
                            fmt.Println(f)
                  }
          }
}
func main() {
          util.PrepareErrorMessages("cleanSeq")
          var optL = flag.Int("l", 150, "maximum length of internal run of Ns")
          var optM = flag.Int("m", 100, "minimum fragment length")
          var optV = flag.Bool("v", false, "print version & " +
                    "program information")
          u := "cleanSeq [-h] [option]... [file]..."
          p := "Cut runs of N from the sequences returned by fur."
          e := "cleanSeq foo.fasta"
          clio.Usage(u, p, e);
          flag.Parse()
          if *optV {
                    util.PrintInfo("cleanSeq")
          }
          files := flag.Args()
          clio.ParseFiles(files, parse, *optL, *optM)
}
