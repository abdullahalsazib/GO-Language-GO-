package main

import "fmt"

func main() {
    fmt.Println("Welcome in Go - loop")

    for i := 0; i < 10; i++ {
        fmt.Println("Hello, World! ", i + 1)
    }

    fmt.Println("[--------like While loop-----------]")
    routCounter := 1
    for routCounter <= 10 {
        if routCounter == 5 {
            goto loc
        }
        if routCounter == 3 + 1 {
            routCounter++
            continue
        }
        fmt.Println("Hello, World", routCounter)
        routCounter++
    }
    loc:
        fmt.Println("Goto jumpin this code..!")
}
