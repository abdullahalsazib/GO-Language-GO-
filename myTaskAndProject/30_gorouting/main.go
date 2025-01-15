package main

import (
	"fmt"
	"time"
)

func prinNumber() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(250 * time.Millisecond)
	}
}
func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func main() {
	//prinNumber()
	go printLetters()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("hello, world")
}
