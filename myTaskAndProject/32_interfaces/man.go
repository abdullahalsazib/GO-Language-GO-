package main

import (
	"fmt"
	"log"
)

type Animaler interface {
	Says() string
	NumberOfLegs() int
}

type Doger struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	fmt.Println("interface")

	dog := Doger{
		Name:  "Samsom",
		Breed: "German shepherd",
	}

	PrintInfo(dog)

	

}

func (d Doger) Says() string {
	return "woof"
}

func (d Doger) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animaler) {
	log.Println("This animal says ", a.Says(), a.NumberOfLegs())
}
