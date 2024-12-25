package main

import (
	"fmt"
	
)

func main() {
    fmt.Println("Welcome in struct: ")
    // not inheritance in golang; No super or parent

    jack := User{"Jack", "jack@main.com", true, 19}

    fmt.Println(jack)
    fmt.Printf("jack details: %+v\n", jack)

    fmt.Printf("Name is: %v and Email is: %v\n", jack.Name, jack.Email)
}

type User struct {
    Name string
    Email string
    Status bool
    Age int 
}
