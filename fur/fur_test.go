package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestFur(t *testing.T) {
	var tests []*exec.Cmd
	d := "test.db"
	test := exec.Command("./fur", "-d", d)
	tests = append(tests, test)
	test = exec.Command("./fur", "-d", d, "-q", "0.5")
	tests = append(tests, test)
	test = exec.Command("./fur", "-d", d, "-q", "0.5", "-w", "150")
	tests = append(tests, test)
	test = exec.Command("./fur", "-d", d, "-q", "0.5", "-w", "150",
		"-t", "8")
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
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
