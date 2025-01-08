package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}

func main() {
	fmt.Println("Hello, channel")

	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}
