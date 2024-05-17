package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestStream(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.fasta"
	test := exec.Command("./stream", f)
	fmt.Println(test)
	tests = append(tests, test)
	test = exec.Command("./stream", "-i", f)
	tests = append(tests, test)
	test = exec.Command("./stream", "-S", "3", "-s", "0.5", f)
	tests = append(tests, test)
	test = exec.Command("./stream", "-S", "3", "-s", "0.5",
		"-m", "0.02", f)
	tests = append(tests, test)
	test = exec.Command("./stream", "-S", "3", "-s", "0.5",
		"-m", "0.02", "-i", f)
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
