package main

import "fmt"

func main() {
	var i int = 23
	var f float32 = float32(i)
	var u uint = uint(f)
	

	fmt.Printf("Int: %v\n", i)
	fmt.Printf("Float: %.3f\n", f)
	fmt.Printf("Uint: %v", u)


}