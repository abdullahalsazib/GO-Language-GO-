package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	fmt.Println("welcome in go struct")

	p1 := Person{Name: "jack", Age: 23, Country: "London"}

	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Country)
}
