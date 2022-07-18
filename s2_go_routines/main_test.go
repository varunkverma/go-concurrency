package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout // saving stdout before running test

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printSomething("testing", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "testing") {
		t.Errorf("Expected to find 'testing' but its not there")
	}
}
