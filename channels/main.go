package main

import "fmt"

// import "time"
import "strings"

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	// time.Sleep(10 * time.Second)
	fmt.Println("Type something,and press enter(enter q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")
		// get user input

		var userInput string

		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		} else {
			ping <- userInput

			// wait for a response

			response := <-pong

			fmt.Println("Response:", response)
		}
	}

	fmt.Println("All done, Closing channels.")
	close(ping)
	close(pong)
}
