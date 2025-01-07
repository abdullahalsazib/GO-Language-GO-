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

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment is usnig fake: ", amount)
}

func main() {
	fmt.Println("Wellcome in go-interface")
	// stripPaymentGw := strip{}
	// razerPaymentGw := rezarpay{}
	fakePayGw := fakePayment{}
	newPayment := payment{
		gateWay: fakePayGw,
	}
	newPayment.makePayment(200.3)	

}
