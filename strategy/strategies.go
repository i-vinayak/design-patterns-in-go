package main

import "fmt"

// Credit Card
type CreditCardStrategy struct {
	Name   string
	Number string
	CVV    string
}

func (cc *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card (%s)", amount, cc.Name)
}

// Pay pal
type PayPalStrategy struct {
	Email string
}

func (pp *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (%s)", amount, pp.Email)
}
