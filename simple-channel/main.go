package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping

		pong <- strings.ToUpper(fmt.Sprintf("%s!!!", s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Enter something to shout and press Enter. Press q to quit.")
	for {
		fmt.Print("-> ")
		var input string
		_, _ = fmt.Scanln(&input)

		if strings.ToLower(input) == "q" {
			break
		}

		ping <- input

		output := <-pong

		fmt.Println("Response:", output)
	}
	fmt.Println("Closing channels")
	close(ping)
	close(pong)
}
