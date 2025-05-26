package main

type PaymentContext interface {
	SetStrategy(PaymentStrategy)
	Pay(float64) string
}

type paymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) PaymentContext {
	return &paymentContext{strategy}
}

func (c *paymentContext) SetStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *paymentContext) Pay(amount float64) string {
	return c.strategy.Pay(amount)
}
