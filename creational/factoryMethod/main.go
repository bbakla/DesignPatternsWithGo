package main

import "fmt"

func main() {
	paymentFactory := new(PaymentFactory)

	payWithCreditCard, _ := paymentFactory.GetPaymentMethod(1)
	payWithCreditCard.Pay(50)
	fmt.Printf("Expense is %f\n", payWithCreditCard.GetExpenses())

	payWithDebitCard, _ := paymentFactory.GetPaymentMethod(2)
	payWithDebitCard.Pay(70)
	fmt.Printf("Expense is %f\n", payWithDebitCard.GetExpenses())

}
