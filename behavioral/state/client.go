package main

import (
	"fmt"
	"log"
)

func main() {
	var name ItemName
	name = "chocolate"
	vendingMachine := NewVendingMachine(name, 20, 3)
	err := vendingMachine.requestItem(name)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(name, 10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem(name)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(name, 2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem(name)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(name, 10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem(name)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
