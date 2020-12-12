package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float32) string
	GetExpenses() float32
}

type DebitCardPayment struct {
	expenses float32
}

func (db *DebitCardPayment) Pay(amount float32) string {
	db.expenses += amount
	return fmt.Sprintf("paying %f with debit card", amount)
}

func (db *DebitCardPayment) GetExpenses() float32 {
	return db.expenses
}

type CreditCardPayment struct {
	expenses float32
}

func (cc *CreditCardPayment) Pay(amount float32) string {
	cc.expenses += amount
	return fmt.Sprintf("paying %f with credit card", amount)
}

func (cc *CreditCardPayment) GetExpenses() float32 {
	return cc.expenses
}
