package mypack

import "fmt"

func Hello(name string, a,b int) int {
	fmt.Println("Hello this is mypack")
	Bye(name)
	return Sum(a, b)
}

func Bye(name string) {
	fmt.Printf("Bye Bye mr %v\n", name)
}

func Sum(a,b int) int {
	return a + b
}