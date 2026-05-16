package main

import "sync"
import "fmt"

var msg string
var wg sync.WaitGroup

func UpdateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "hello world"

	wg.Add(2)
	go UpdateMessage("Hello, universe!")
	go UpdateMessage("Hello, cosmos!")

	wg.Wait()

	fmt.Println(msg)
}
