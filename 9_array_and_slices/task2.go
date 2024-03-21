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

func main() {
	// Write your code here.
	numbers := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	sum := NestedSliceSum(numbers)
	for _, s := range sum {
		println(s)
	}
}
