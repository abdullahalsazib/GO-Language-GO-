package main

import "fmt"

// interface -->
type paymentr interface {
	pay(amount float32)
}

type payment struct {
	gateway paymentr
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe: ", amount)
// }

// type fakePay struct{}

// func (f fakePay) pay(amount float32) {
// 	fmt.Println("making fake payment using fake payment method", amount)
// }

// paypal
type paypal struct{}

func (pa payment) pay(amount float32) {
	fmt.Println("Making payment using paypal..!", amount)
}
func main() {
	fmt.Println("Welcome for Go-Lang Interface..!")
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	paypalPaymentGw := paypal{}

	newPayment := payment{
		gateway: paypalPaymentGw,
	}

	newPayment.makePayment(100)
}
