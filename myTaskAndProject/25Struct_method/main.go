package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and age is %d. i'm from %s", p.Name, p.Age, p.Country)
}

func (p *Person) CelebrateBirthday() {
	p.Age++
	fmt.Printf("Happy Birthday, %s! You are now %d years old.\n", p.Name, p.Age)
}

func main() {
	fmt.Println("hello, world")
	p1 := Person{Name: "jack", Age: 23, Country: "Bangladesh"}

	p1.Greet()
	fmt.Println("")
	p1.CelebrateBirthday()
}
