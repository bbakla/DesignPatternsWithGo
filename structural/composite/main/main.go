package main

import (
	"fmt"
)

type doer interface {
	doIt()
}

type ZeroInit struct {
}

func (zero *ZeroInit) doIt() {
	fmt.Println("doing it")
}

type Composition struct {
	zeroInit ZeroInit //direct usage
	ZeroInit          //embedded usage
	Closure  func(string)
}

type CompositionViaInterface struct {
	doer
}

func main() {
	composition := Composition{
		Closure: ClosureFunc,
	}

	composition.Closure("I am ")
	composition.zeroInit.doIt() //direct usage
	composition.doIt()          //embedded usage

	compositionviaInterface := CompositionViaInterface{
		&ZeroInit{},
	}

	compositionviaInterface.doIt()

}

func ClosureFunc(who string) {
	fmt.Printf("%s in closure func\n", who)
}
