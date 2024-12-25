package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	fmt.Println("Welcome -> in Go Defer <| => ")
	hello()
}

func hello() {
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
}