package main

import "fmt"

func main() {
	// Declaring and init an Array -->
	array := []int{1, 2, 3, 4, 5}
	fmt.Println("array: ", array)

	// Accessing an index of array element >-->
	fmt.Printf("Element of index 1 is : %v\n", array[1])

	// Modifyied the Array Element using the index like 1 no index
	array[1] = 10 // array[1] = 2 --> array[1] = 10 <--> O(n)
	fmt.Printf("Array[1] = %v\n", array[1])

	// inserting an element at position 2

	index := 2
	value := 7

	// array[index] = value

	// also use the append() <--> O(n)
	array = append(array[:index], append([]int{value}, array[index:]...)...)
	fmt.Printf("Now the array is : %v\n", array)

	// array size
	fmt.Printf("Array size: %v\n", len(array))

	// Delete element using the append()
	index = 2
	array = append(array[:index], array[index+1:]...)
	fmt.Printf("After delete at index %v and now Array is %v\n", index, array)
}
