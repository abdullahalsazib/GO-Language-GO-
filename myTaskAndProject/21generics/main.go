package main

import "fmt"

// also use in [T any] for any interface
func printSlice[T int | string | bool ](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct{
	element []T 
}

func main() {
	fmt.Println("Wellcome in go Generics")
	// num := []int{1, 2, 3, 4, 5}
	name := []string{"jack", "sparrow", "sazib", "Abdullah"}
	printSlice(name)

	myStack := stack[string]{
		element: name, 
	}

	fmt.Println(myStack)
}
