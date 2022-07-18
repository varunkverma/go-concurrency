package main

import (
	"fmt"
	"sync"
	"time"
)

// constants
const hunger = 3

// variables - Philosophers
var philosophers = []string{
	"Plato",
	"Socarets",
	"Aristotle",
	"Pascal",
	"Locke",
}

var wg sync.WaitGroup

var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second

var finishSequence = []string{}

func diningProblem(philosopher string, leftFork, rightFork, finMutex *sync.Mutex) {
	defer wg.Done()
	// print a message
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry")
		time.Sleep(sleepTime)
		// lock both forks mutexes
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork of his left\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork of his right\n\n", philosopher)
		// print a message
		fmt.Println(philosopher, "has both forks, and is eating")
		time.Sleep(eatTime)

		// giving the philosopher time to think
		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)
		// unlock the fork mutexes
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on this right\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on this left\n\n", philosopher)
		time.Sleep(sleepTime)
	}

	// print out done message
	fmt.Println(philosopher, "is satisfied")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table")
	finMutex.Lock()
	finishSequence = append(finishSequence, philosopher)
	finMutex.Unlock()

}

func main() {
	// print intro
	fmt.Println("The Dining Philosopher Problem")
	fmt.Println("_______________________________")

	wg.Add(len(philosophers))

	// we need to create a mutex for the very first fork(the one to the left of the first philosopher). We create it as a pointer, since a sync.Mutex must not be copied adter its initial use.
	forkLeft := &sync.Mutex{}

	finMutex := &sync.Mutex{}
	// spawn one go routine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		forkRight := &sync.Mutex{}

		// call a goRoutine
		go diningProblem(philosophers[i], forkLeft, forkRight, finMutex)

		// create the next philosopher's left fork (which is the current philosopher's right fork). Note that we are not coping a mutex here; we are making forkLeft equal to the pointer to an existing mutex, so it points to the same location in memory, and does not copy it.
		forkLeft = forkRight
	}

	wg.Wait()
	fmt.Println("_______________________________")
	fmt.Println("The Table is empty")
	fmt.Println("Finishing sequence of philosophers leaving the table:")
	for _, philosopher := range finishSequence {
		fmt.Printf("\t%s\n", philosopher)
	}
}
