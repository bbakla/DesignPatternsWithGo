package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem(name ItemName) error {
	if i.vendingMachine.items[name].count == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) addItem(name ItemName, count int) error {
	fmt.Printf("%d items added to %s\n", count, name)
	i.vendingMachine.incrementItemCount(name, count)
	return nil
}

func (i *hasItemState) insertMoney(name ItemName, money int) error {
	return fmt.Errorf("please select item first")
}

func (i *hasItemState) dispenseItem(name ItemName) error {
	return fmt.Errorf("please select item first")
}
