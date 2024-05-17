package main

import (
          "github.com/evolbioinf/fur/util"
          "github.com/evolbioinf/clio"
          "flag"
          "text/tabwriter"
          "os"
          "github.com/evolbioinf/sus"
          "fmt"
)

func main() {
          util.PrepareErrorMessages("madis")
          u := "madis [option]..."
          p := "Calculate the match length distribution."
          e := "madis -l 90000 -g 0.45"
          clio.Usage(u, p, e)
          optL := flag.Int("l", 100000, "sequence length")
          optG := flag.Float64("g", 0.5, "GC content")
          msg := "quantile (default probability density)"
          optQ := flag.Float64("q", 0.0, msg)
          optV := flag.Bool("v", false, "version")
          flag.Parse()
          if *optV {
                    util.PrintInfo("madis")
          }
          w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
          if *optQ > 0 {
                    fq := sus.Quantile(*optL, *optG, *optQ) - 1
                    fmt.Fprintf(w, "#p\tQ(p)\t\n")
                    fmt.Fprintf(w, "%g\t%d\t\n", *optQ, fq)
          } else {
                    fmt.Fprintf(w, "#x\tPDF(x)\tCDF(x)\t\n")
                    x := 0
                    cdf := 0.0
                    pdf := sus.Prob(*optL, *optG, x)
                    for pdf == 0 {
                              x++
                              pdf = sus.Prob(*optL, *optG, x)
                    }
                    for pdf > 0 {
                              cdf += pdf
                              fmt.Fprintf(w, "%2d\t%.3g\t%.3g\t\n", x-1, pdf, cdf)
                              x++
                              pdf = sus.Prob(*optL, *optG, x)
                    }
          }
          w.Flush()
}
