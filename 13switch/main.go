package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    fmt.Println("Wellcome... in Switch Case")
    
    rand.Seed(time.Now().UnixNano())
    diceNumber := rand.Intn(3) + 1

    fmt.Println("Vaue of dice is", diceNumber)

    switch diceNumber {
    case 1:
        fmt.Println("DiceNumber is 1")
    case 2:
        fmt.Println("DiceNumber is 2")
        fallthrough
    case 3:
        fmt.Println("DiceNumber is 3")
    case 4:
        fmt.Println("DiceNumber is 4")
    }
}
