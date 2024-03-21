package main

import "fmt"

func ExampleFunction(x string, y int) (string, int) {
	return x, y
}

func main() {
	// Print line
	fmt.Println("Hello, World!", 123)

	// Print formatted
	fmt.Printf("Decimal:%d, Binary:%b, Scientific:%e\n", 10, 10, 10.0)

	// Print string and float
	fmt.Printf("String:%s, Float:%.2f\n", "Hello, World!", 10.0)

	// Print integer with padding
	fmt.Printf("Padded integer:%05d\n", 10)

	// Escape characters
	fmt.Printf("Escaped characters\"%%1\"\n")

	// String formatting
	s := fmt.Sprintf("Hello, World!")
	fmt.Println(s)

	// Function call
	x, y := ExampleFunction("Hello, World!", 123)
	fmt.Println(x, y)
}
