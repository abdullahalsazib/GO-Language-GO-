package main

import "fmt"
func main()  {
   fmt.Println("Welcome to array in golang")

    var fruitList = [3]string{"Apple", "Orange", "Lemon"}

    temp := fruitList[2]
    fruitList[2] = "Tomato"

    fmt.Println(temp)
    fmt.Println(fruitList)
    fmt.Println(len(fruitList))
}
