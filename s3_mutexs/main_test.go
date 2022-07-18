package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	oldStdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)

	os.Stdout = oldStdOut

	if !strings.Contains(output, "34320.00") {
		t.Error("wrong balance returned")
	}
}
