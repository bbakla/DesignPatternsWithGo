package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem(name ItemName) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) addItem(name ItemName, count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) insertMoney(name ItemName, money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *hasMoneyState) dispenseItem(name ItemName) error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.items[name].count -= 1
	if i.vendingMachine.items[name].count == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
