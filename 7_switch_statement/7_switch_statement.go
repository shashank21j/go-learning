package main

import "fmt"

func main() {
	switch a := 10; {
	case a == 10:
		fmt.Println("a is 10")
	case a > 10:
		fmt.Println("a is greater than 10")
	default:
		fmt.Println("a is less than 10")
	}

	switch b := 0; {
	case b <= 20:
		fmt.Println("b is less than or equal to 20")
		fallthrough
	case b < 10:
		fmt.Println("b is less than 10")
		fallthrough
	default:
		fmt.Println("default case")
	}

	c := 'h'
	switch c {
	case 'h', 'i', 'j':
		fmt.Println("c is h")
	case 'a', 'b', 'c':
		fmt.Println("c is a")
	default:
		fmt.Println("default case")
	}
}
