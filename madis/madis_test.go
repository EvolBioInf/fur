package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestMadis(t *testing.T) {
	var tests []*exec.Cmd
	test := exec.Command("./madis", "-l", "1000000", "-g", "0.1")
	tests = append(tests, test)
	test = exec.Command("./madis", "-l", "1000000", "-g", "0.1",
		"-q", "0.05")
	tests = append(tests, test)
	for i, test := range tests {
		g, e := test.Output()
		if e != nil {
			t.Error(e)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		w, e := os.ReadFile(f)
		if e != nil {
			t.Error(e)
		}
		if !bytes.Equal(g, w) {
			t.Errorf("get:\n%s\nwant:\n%s\n", g, w)
		}
	}
}
