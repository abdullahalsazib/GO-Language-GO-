package main

import "fmt"

func oneDimentionalArry(n int) []int {
	var a int
	// Array
	var myArray = []int{}

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		myArray = append(myArray, a)
	}

	return myArray
}

// Binary Search .._>
func binarySearch(arr []int, target int) int{
	var left, right = 0, len(arr) - 1

	for left <= right {
		mid := left + (right + left) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid- 1
		}
	}

	return -1
}

func main() {
	var _n, _target int 
	fmt.Printf("Enter your array size: ")
	fmt.Scanf("%d", &_n)

	fmt.Printf("Enter your %v size array: \n", _n)
	var _arr = oneDimentionalArry(_n)

	fmt.Printf("Enter your target value: ")
	fmt.Scanf("%d", &_target)

	
	var idx = binarySearch(_arr, _target)

	if idx != -1 {
		fmt.Printf("your array index %v is found.. ! %v", idx, _arr[idx])
	}else {
		fmt.Println("Your result is not found")
	}

    fmt.Println("Hello World")

}
