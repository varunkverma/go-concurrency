package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	defer timer("main")()
	// varibale for bank balance
	var bankBalance int
	// print out the starting value
	fmt.Printf("Initial account balance: %d.00\n", bankBalance)
	// define weekly value revenue
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 500,
		},
		{
			Source: "Gifts",
			Amount: 10,
		},
		{
			Source: "Part Time Job",
			Amount: 50,
		},
		{
			Source: "Investments",
			Amount: 100,
		},
	}

	wg.Add(len(incomes))

	var mutex sync.Mutex

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income, m *sync.Mutex) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				m.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Printf("On week %d, you earned %d.00 from %s, current bank balance: %d.00\n", week, income.Amount, income.Source, bankBalance)
				m.Unlock()
			}
		}(i, income, &mutex)
	}
	wg.Wait()
	// print out final balance
	fmt.Printf("Final bank balance: %d.00\n", bankBalance)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v seconds\n", name, time.Since(start).Seconds())
	}
}

// var msg string

// var wg sync.WaitGroup

// // func updateMessage(s string, m *sync.Mutex) {
// // 	defer wg.Done()
// // 	m.Lock()
// // 	msg = s
// // 	m.Unlock()
// // }

// // func main() {
// // 	msg = "Hello world"

// // 	var mutex sync.Mutex

// // 	wg.Add(2)
// // 	go updateMessage("Ola World", &mutex)
// // 	go updateMessage("Namaste Duniya", &mutex)
// // 	wg.Wait()

// // 	fmt.Println(msg)
// // }
