package main

import "fmt"

func multiply(a int, b int, ch chan int) {
	ch <- a * b
}
func addition(a, b int, ch chan int) {
	ch <- a + b
}
func main() {
	ch := make(chan int)
	go multiply(2, 3, ch)

	res := <-ch
	// res := multiply(2, 3)
	fmt.Println("multiply: ", res)

	// close the ch
	close(ch)

	go addition(10, 20, ch)
	res = <-ch
	fmt.Println(res)
}
