package main

import (
	"errors"
	"fmt"
)


var ErrDividByZero = errors.New("divide by zero") 
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDividByZero
    }
    return a / b, nil
}
func main() {
    fmt.Println("Error in Go-lang")
    
    a, b := 10, 3
    result, err := Divide(a, b)
    if err != nil {
        switch {
        case errors.Is(err, ErrDividByZero):
            fmt.Println("divide by zero")
        default:
            fmt.Println("unexpected divide")
        }
        return
    }
    fmt.Printf("%d / %d = %d\n", a, b, result)
    
}
