package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Microsecond)
	}
}

func main() {
	fmt.Println("_Concurrency_gorouting")
	printNumber()
	fmt.Println("main function running....")

	time.Sleep(3 * time.Second)
	fmt.Println("Main function exit")
}
