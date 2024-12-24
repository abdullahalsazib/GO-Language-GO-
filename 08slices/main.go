package main

import (
	"fmt"
	"sort"
)

func main(){
    fmt.Println("Welcome to array on slices")

    var fruitList = []string{"Apple", "Tomato", "Orange"}

    fruitList = append(fruitList, "Lemon")
    fruitList = append(fruitList, "Mango", "Banana")
    fmt.Println(fruitList)

   // fruitList = append(fruitList[1:)])

    fmt.Printf("Type of fruitlist is %T\n", fruitList)
    fmt.Println(fruitList)

    highScore := make([]int, 4)

    highScore[0] = 234
    highScore[1] = 345
    highScore[2] = 456
    highScore[3] = 678

    highScore = append(highScore, 333,445,666,000)


    fmt.Println(sort.IntsAreSorted(highScore))

    sort.Ints(highScore)
    fmt.Println(highScore)

    fmt.Println(sort.IntsAreSorted(highScore))


    // how to remove a value from slices based on index
    var courses = []string{"javascript", "Nodejs", "reactjs", "Python", "ruby", "java"}

    fmt.Println(courses)
    var index int = 2 
    courses = append(courses[:index], courses[index+1:]...)
    
    fmt.Println(courses)








}
