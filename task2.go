package main

func NestedSliceSum(numbers [][]int) []int {
	// Write your code here.
	sum := []int{}
	for _, slice := range numbers {
		s := 0
		for _, num := range slice {
			s += num
		}
		sum = append(sum, s)
	}
	return sum
}
