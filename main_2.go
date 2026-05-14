package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var msg string

var wg sync.WaitGroup


func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func Test_upateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("epsilon")
	wg.Wait()
	if msg != "epsilon" {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r,w,_ := os.Pipe()
	os.Stdout = w
	msg = "epsilon"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Error("Expected to find epsilon, but it is not ther")
	}

}

func Test_demoMain(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	demoMain()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe, but it is not ther)")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos, but it is not ther)")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, word, but it is not ther)")
	}
}

func demoMain() {
	msg = "Hello, world"
	wg.Add(1)

	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()
}
