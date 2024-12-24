package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Enter your age: ")
	fmt.Scanf("%d\n", &age)


	if age >= 0 && age <= 17{
		fmt.Println("You'r Child!")
	} else if age >= 18 && age <= 30 {
		fmt.Println("You'r Adult!")
	} else if age >= 31 && age <= 60 {
		fmt.Println("Your'r Young!")
	} else if age >= 61 && age <= 90 {
		fmt.Println("You'r Old!")
	} else if age < 0 || age <= -1 {
		fmt.Println("Invalid.. your age is not valid!")
	}
}