package main

import "fmt"

type Spacker interface {
	Spack() string
}

type Persone struct {
	Name string
}

// implement the speak method for person type
func (p Persone) Spack() string {
	return "Hello, my name is " + p.Name
}

type Robot struct {
	Model string
}

func (r Robot) Spack() string {
	return "I am a robot of model " + r.Model
}

func Introduce(s Spacker) {
	fmt.Println(s.Spack())
}

func main() {
	fmt.Println("hello, world")

	p := Persone{Name: "jack"}
	r := Robot{Model: "22l20033k"}

	Introduce(p)
	Introduce(r)
}
