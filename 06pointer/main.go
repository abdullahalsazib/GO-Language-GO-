package main

import "fmt"

func main() {
    fmt.Println("Welcome to a class on pointers")
    var a int = 10
    // Pointer 
    //var p *int
    myNumber := 23
    var ptr = &myNumber
    
    *ptr += 2 // *prt = *ptr + 2

    fmt.Println(a, *ptr)
}
