package main

type state interface {
	addItem(name ItemName, count int) error
	requestItem(name ItemName) error
	insertMoney(name ItemName, money int) error
	dispenseItem(name ItemName) error
}
