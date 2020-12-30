package main

import (
	"fmt"
)

type Component interface {
	doIt()
}

type Leaf struct {
}

func (zero *Leaf) doIt() {
	fmt.Println("doing it")
}

type Composition struct {
	leaf    Leaf //direct usage
	Leaf         //embedded usage
	Closure func(string)
}

type CompositionViaInterface struct {
	Component
}

func main() {
	composition := Composition{
		Closure: ClosureFunc,
	}

	composition.Closure("I am ")
	composition.leaf.doIt() //direct usage
	composition.doIt()      //embedded usage

	compositionViaInterface := CompositionViaInterface{
		&Leaf{},
	}

	compositionViaInterface.doIt()

}

func ClosureFunc(who string) {
	fmt.Printf("%s in closure func\n", who)
}
