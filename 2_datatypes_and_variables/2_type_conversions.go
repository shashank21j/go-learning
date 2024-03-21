package main

import "fmt"

func main() {
	var x int = 10
	fmt.Printf("%T %v\n", x, x)

	var y float64 = float64(x)
	fmt.Printf("%T %v\n", y, y)
}
