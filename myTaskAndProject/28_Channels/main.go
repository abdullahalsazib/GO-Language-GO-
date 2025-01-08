package main

import "fmt"

func worker(ch chan int) {
	fmt.Println("Worker started..")
	ch <- 42
	fmt.Println("worker finished..")
}

func main() {
	fmt.Println("Channels")
	ch := make(chan int) // create a channel of type int

	go worker(ch)

	result := <-ch
	fmt.Println(result)
}
