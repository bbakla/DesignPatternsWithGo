package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	goroutines()

	//We give this sleep to wait until goroutine that runs helloWorld to finish before our main function ends.
	time.Sleep(time.Second)
}

func goroutines() {
	go helloWorld()

	go func(message string) {
		println("Hello from anonymous functions")
		println(message)
	}("message")

	messagePrinter := func(message string) {
		println(message)
	}

	go messagePrinter("AAA")
	go messagePrinter("BBBBB")

}

func waitgroups() {
	var wait sync.WaitGroup
	wait.Add(1)

	messagePrinter := func(message string) {
		println(message)
		wait.Done()
		//wait.Add(-1) // That is also possible.
	}

	go messagePrinter("HELLOOO")

	//Lets add multiple goroutines

	wait.Add(5)

	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		go messagePrinter(s)
	}

	wait.Wait()
}

func helloWorld() {
	fmt.Println("Hello World")
}
