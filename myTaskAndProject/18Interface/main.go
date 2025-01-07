package main

import "fmt"

type paymentr interface {
	pay(amount float32)
}

type payment struct {
	gateWay paymentr
}

func (p payment) makePayment(amount float32) {
	// razarPaygw := strip{}
	p.gateWay.pay(amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making paymenet is using paypal: ", amount)
}

func main() {
	fmt.Println("Wellcome in go-interface")
	// stripPaymentGw := strip{}
	// razerPaymentGw := rezarpay{}
	// fakePayGw := fakePayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateWay: paypalGw,
	}
	newPayment.makePayment(200.10)

}
