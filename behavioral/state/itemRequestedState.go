package main

import "fmt"

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem(name ItemName) error {
	return fmt.Errorf("Item already requested")
}

func (i *itemRequestedState) addItem(name ItemName, count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *itemRequestedState) insertMoney(name ItemName, money int) error {
	if money < i.vendingMachine.items[name].price {
		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.items[name].price)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)

	return nil
}
func (i *itemRequestedState) dispenseItem(name ItemName) error {
	return fmt.Errorf("Please insert money first")
}
