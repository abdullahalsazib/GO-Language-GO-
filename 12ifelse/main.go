package main
import "fmt"

func main() {
    fmt.Println("if else in golang")
    var result string

    loginCount := 23
    if loginCount < 10 {
        result = "Regular user"
    }else if loginCount > 10 {
        result = "Watch out"
    } else {
        result = "Not Regular user"
    }

    fmt.Println(result)

    if num := 3; num < 10 {
        fmt.Println("Num is less then 10")
    } else {
        fmt.Println("Num is Not less then 10")
    }

    
}
