// The package util provides utility functions for the fur package.
package util

import (
	"fmt"
	"github.com/evolbioinf/clio"
	"log"
	"os"
	"strings"
)

var version, date string

// PrintInfo takes as arguments the name. As output it prints this name, together with the version, compilation date, authors' names and email address, and the program's license. Then it exits.
func PrintInfo(n string) {
	v := version
	d := date
	a := "Bernhard Haubold,Beatriz Vieira Mourato," +
		"Ivan Tsers"
	e := "haubold@evolbio.mpg.de,mourato@evolbio.mpg.de," +
		"tsers@evolbio.mpg.de"
	l := "Gnu General Public License, " +
		"https://www.gnu.org/licenses/gpl.html"
	clio.PrintInfo(n, v, d, a, e, l)
	os.Exit(0)
}

// Check takes as argument an error, prints that error, and exits.
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// PrepareErrorMessages takes as argument the program name and sets this as the prefix for error messages from the log package.
func PrepareErrorMessages(name string) {
	m := fmt.Sprintf("%s - ", name)
	log.SetPrefix(m)
	log.SetFlags(0)
}

// IsFasta takes as argument the name of a file and determines whether or not it has the extension of a FASTA file.
func IsFasta(f string) bool {
	a := strings.Split(f, ".")
	s := a[len(a)-1]
	if s == "fasta" || s == "fna" || s == "ffn" ||
		s == "faa" || s == "frn" || s == "fa" {
		return true
	}
	return false
}

// Version returns the program version.
func Version() string {
	return version
}
