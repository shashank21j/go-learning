package main

import "fmt"

func main() {
	// Create a slice with initial elements
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	// Append an element to the slice
	numbers = append(numbers, 6)
	fmt.Println("Slice after appending an element:", numbers)

	// Create a slice with make function, specifying length and capacity
	// make([]Type, length, capacity)
	moreNumbers := make([]int, 3, 5)
	moreNumbers[0] = 7
	moreNumbers[1] = 8
	moreNumbers[2] = 9
	fmt.Println("Slice created with make:", moreNumbers)

	// Slice a slice to create a new slice
	subSlice := numbers[1:4] // This will include elements at index 1, 2, and 3
	fmt.Println("Sub-slice of original slice:", subSlice)
}
