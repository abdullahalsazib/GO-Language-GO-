package main
import "fmt"


func sayHello() {
    fmt.Println("Hello from the Goroutines")
}
func main() {
    go sayHello()
    fmt.Println("Hello from the main Function")


    var input string
    fmt.Scanln(&input)

    fmt.Println("Your string is: ", input)
}
