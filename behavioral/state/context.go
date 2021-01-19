package main

import "fmt"

type ItemName string
type Item struct {
	count int
	price int
}

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	//itemCount int
	//itemPrice int
	items map[ItemName]*Item
}

func NewVendingMachine(itemName ItemName, itemCount int, itemPrice int) *vendingMachine {
	items := make(map[ItemName]*Item)

	items[itemName] = &Item{
		count: itemCount,
		price: itemPrice,
	}

	v := &vendingMachine{
		//itemCount: itemCount,
		//itemPrice: itemPrice,
		items: items,
	}

	hasItemState := &hasItemState{v}
	itemRequestedState := &itemRequestedState{v}
	hasMoneyState := &hasMoneyState{v}
	noItemState := &noItemState{v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	return v
}

func (v *vendingMachine) requestItem(name ItemName) error {
	return v.currentState.requestItem(name)
}

func (v *vendingMachine) addItem(name ItemName, count int) error {
	return v.currentState.addItem(name, count)
}

func (v *vendingMachine) insertMoney(name ItemName, money int) error {
	return v.currentState.insertMoney(name, money)
}

func (v *vendingMachine) dispenseItem(name ItemName) error {
	return v.currentState.dispenseItem(name)
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(itemName ItemName, count int) {
	fmt.Printf("Adding %d items\n", count)
	item := v.items[itemName]
	item.count += count
}
