package main

import (
	"fmt"
	"math"
)

func main() {
	var sr string = "Hello, World!"
	const PI float32 = math.Pi

	fmt.Printf("%v\t %T", sr, sr)
	fmt.Printf("\nPI value is: %.4f\n", PI)
}