package main

import (
	"errors"
	"fmt"
)

type Person struct {
    Name string
    Age int
}
func main() {
    fmt.Println("Hello, World!")

    // Array and Slice 
    // Slice in Dynamic-size
    slice := []int{1,2,3}
    slice = append(slice, 3)
    fmt.Println("Slice is: ", slice)
    p := Person{"Jack Sparrow", 29}
    fmt.Printf("Name: %s\nAge: %v", p.Name, p.Age)
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error", err)
    } else {
        fmt.Println("Result: ", result)
    }
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Divide by zero")
    } 
    return a / b, nil
}
