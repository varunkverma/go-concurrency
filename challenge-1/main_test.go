package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	oldValue := msg

	testString := "Test String"

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage(testString, &wg)
	wg.Wait()

	if msg != testString {
		t.Errorf("expected msg value to be %s, but got %s", testString, msg)
	}

	msg = oldValue
}

func Test_printMessage(t *testing.T) {
	oldStdOut := os.Stdout
	oldMsg := msg

	msg = "Test"
	r, w, _ := os.Pipe()
	os.Stdout = w

	printMessage()

	w.Close()

	outputInBytes, _ := io.ReadAll(r)
	output := string(outputInBytes)

	if !strings.Contains(output, msg) {
		t.Errorf("expected stdout to contain %s, but it contains %s", msg, output)
	}

	os.Stdout = oldStdOut
	msg = oldMsg
}

func Test_main(t *testing.T) {
	var wg sync.WaitGroup
	oldStdOut := os.Stdout
	oldMsg := msg

	r, w, _ := os.Pipe()
	os.Stdout = w

	testString := "Hello, universe!"

	wg.Add(1)
	go updateMessage(testString, &wg)
	wg.Wait()

	if msg != testString {
		t.Errorf("expected msg value to be %s, but got %s", testString, msg)
	}

	printMessage()

	w.Close()

	outputInBytes, _ := io.ReadAll(r)
	output := string(outputInBytes)

	if !strings.Contains(output, msg) {
		t.Errorf("expected stdout to contain %s, but it contains %s", msg, output)
	}
	os.Stdout = oldStdOut
	msg = oldMsg
}
