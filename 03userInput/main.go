package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome, for user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the text: ")

	input, _ := reader.ReadString('\n')

	fmt.Println(input)
	fmt.Printf("%v%T", input, input)

}