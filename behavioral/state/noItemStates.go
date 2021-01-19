package main

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem(name ItemName) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) addItem(name ItemName, count int) error {
	i.vendingMachine.incrementItemCount(name, count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)

	return nil
}

func (i *noItemState) insertMoney(name ItemName, money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) dispenseItem(name ItemName) error {
	return fmt.Errorf("item out of stock")
}
