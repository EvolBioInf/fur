package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestMakeFurDb(t *testing.T) {
	var tests []*exec.Cmd
	p := "./makeFurDb"
	a := "targets"
	n := "neighbors"
	d := "test.db"
	test := exec.Command(p, "-t", a, "-n", n, "-d", d, "-o")
	tests = append(tests, test)
	r := "t2.fasta"
	test = exec.Command(p, "-t", a, "-n", n, "-d", d, "-o",
		"-r", r)
	tests = append(tests, test)
	test = exec.Command(p, "-t", a, "-n", n, "-d", d, "-o",
		"-r", r, "-T", "1")
	tests = append(tests, test)
	for i, test := range tests {
		_, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		get, err := os.ReadFile(d + "/e.fasta")
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
