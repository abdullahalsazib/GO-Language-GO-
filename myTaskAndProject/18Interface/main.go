package main

import "fmt"

type payment struct {
	gateWay strip
}

func (p payment) makePayment(amount float32) {
	// razarPaygw := strip{}
	p.gateWay.pay(100)
}

type rezarpay struct{}

func (r rezarpay) pay(amount float32) {
	fmt.Println("making payment is using razarpay: ", amount)
}

type strip struct{}

func (s strip) pay(amount float32) {
	fmt.Println("making payment is using strip: ", amount)
}

type fakePayment struct{}
func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment is usnig fake: ", amount)
}

func main() {
	fmt.Println("Wellcome in go-interface")
	stripPaymentGw := strip{}
	newPayment := payment{
		gateWay: stripPaymentGw,
	}
	newPayment.makePayment(100)

}
