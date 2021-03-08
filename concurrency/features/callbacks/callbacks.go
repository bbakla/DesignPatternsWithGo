package main

import (
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main() {

	wait.Add(1)

	toUpperSync("hello", func(string2 string) {
		println(string2)
	})

	toUpperAsync("hello world", func(string2 string) {
		println(string2)
		wait.Done()
	})

	println("Waiting async response..")
	wait.Wait()
}

func toUpperSync(word string, f func(string2 string)) {
	f(strings.ToUpper(word))
}

func toUpperAsync(word string, f func(string2 string)) {
	go f(strings.ToUpper(word))
}
