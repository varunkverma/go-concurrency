package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	errCh := make(chan error)
	p.quit <- errCh
	return <-errCh
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {

		fmt.Printf("Received order number %d!\n", pizzaNumber)

		randomPizzaSuccess := rand.Intn(12) + 1 // random value between from 1 to 12, inclusive
		msg := ""
		success := false

		if randomPizzaSuccess < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		delay := rand.Intn(5) + 1
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if randomPizzaSuccess <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if randomPizzaSuccess <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzaria(pizzaMaker *Producer) {
	// keep track of which pizza are we making
	i := 0

	// run forever or until we receive a quit notification
	// try to make pizza
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data channel of the Producer)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channel
				close(pizzaMaker.data)
				close(quitChan)

				// leave go routine
				return
			}
		}
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzaria is open for business")
	color.Cyan("_________________________________")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzaria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really MAD!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel! ***", err)
			}
		}
	}

	// print out the end message
	color.Cyan("________________")
	color.Cyan("Done for the day")
	color.Cyan("We make %d pizzas, but failetd to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an aweful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!!!")
	default:
		color.Green("It was a great day!")
	}
}
