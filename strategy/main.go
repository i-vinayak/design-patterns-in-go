package main

import "fmt"

func main() {
	// Choose payment method at runtime
	payPal := &PayPalStrategy{Email: "user@example.com"}
	creditCard := &CreditCardStrategy{Name: "John Doe", Number: "1234-5678", CVV: "123"}

	// Use PayPal
	context := NewPaymentContext(payPal)
	fmt.Println(context.Pay(100.0))

	// Switch to Credit Card
	context.SetStrategy(creditCard)
	fmt.Println(context.Pay(250.5))
}
