package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	r := add(10, 20)
	fmt.Println("r = ", r)
}
