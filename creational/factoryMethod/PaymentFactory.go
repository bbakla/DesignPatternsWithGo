package main

import "errors"

const (
	CreditCard = 1
	DebitCard  = 2
)

type PaymentFactory struct {
}

func (payment *PaymentFactory) GetPaymentMethod(paymentType int) (PaymentMethod, error) {
	switch paymentType {
	case 1:
		return &CreditCardPayment{}, nil
	case 2:
		return &DebitCardPayment{}, nil
	default:
		return nil, errors.New("select a proper payment method")
	}
}
