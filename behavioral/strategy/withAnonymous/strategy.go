package main

import "fmt"

func main() {
	strategy1 := func() {
		fmt.Println("First strategy")
	}

	strategy2 := func() {
		fmt.Println("Second strategy")
	}

	context := Context{Strategy: strategy1}
	context.doIt()

	context.Strategy = strategy2
	context.doIt()

}
