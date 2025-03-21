package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Type Assertion")

	var i interface{} = 42

	fmt.Println("i type is ", reflect.TypeOf(i))
	value, ok := i.(int)

	if ok {
		fmt.Println("Converted Value:", value)
	} else {
		fmt.Println("Converted Faild!")
	}
}
