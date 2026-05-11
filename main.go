package main

import "fmt"
import "sync"
//import "time"

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)

}

func main() {
	//go printSomething("This is the first ting to be printed!")

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma", 
		"delta", 
		"epsilon", 
		"zeta", 
		"eta", 
		"theta", 
		"iota",
	}

	//wg.Add(9)
	wg.Add(len(words))

	for i,x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
	//time.Sleep(1 * time.Second)
	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
