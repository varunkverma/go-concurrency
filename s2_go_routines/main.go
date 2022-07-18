package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the WaitGroup
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"apple",
		"bag",
		"cat",
		"dog",
		"eat",
	}

	wg.Add(len(words)) // the number of things you need to wait for

	for i, w := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}
	wg.Wait()

	// time.Sleep(1 * time.Second) // bad way
	wg.Add(1)
	printSomething("Second", &wg)
}
