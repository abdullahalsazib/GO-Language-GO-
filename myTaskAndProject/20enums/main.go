package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Perpared
	Delivered
)

// orderStatus for string

type OrderStatus2 string

const (
	Received2  OrderStatus2 = "received2"
	Confirmed2              = "confirmed2"
	Perpared2               = "perpared2"
	Delivered2              = "delivered2"
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("Change order status to: ", status)
}
func ChangeOrderStatus2(status OrderStatus2) {
	fmt.Println("String change order status: ", status)
}

func main() {
	fmt.Println("wellcome in go enums")
	ChangeOrderStatus(Received)
	fmt.Println("for string emuns")
	ChangeOrderStatus2(Received2)
}
