package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Result is: ", result)
}

func greeter() {

	fmt.Println("Hello, form Go lang")

}
func proAdder(value ...int) int {
	totel := 0

	if 10 >= 5 {
		fmt.Println("Hello World --> hello")
	}

	for _, val := range value {
		totel += val
	}
	return totel

}
