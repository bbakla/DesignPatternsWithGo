package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	originator := &Originator{state: "initialState"}

	careTaker := &Caretaker{
		Originator: originator,
	}

	careTaker.createBackup()
	originator.setState("secondState")
	careTaker.createBackup()
	originator.setState("thirdState")
	fmt.Printf("State is %s\n", originator.state)

	careTaker.undo(0)
	fmt.Printf("State is %s\n", originator.state)
}
