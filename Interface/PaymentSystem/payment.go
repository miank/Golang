package paymentsystem

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type Stripe struct{}

func (s Stripe) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via Stripe \n", amount)
	return nil
}

type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via PayPal \n", amount)
	return nil
}

func MakePayment(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount)
}
