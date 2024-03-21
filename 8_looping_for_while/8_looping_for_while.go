package main

import (
	"fmt"
)

func main() {
	// For loop
	str := "World!"
	for i, ch := range str {
		fmt.Printf("%d %c %T\n", i, ch, ch)
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c %T\n", str[i], str[i])
	}

	// While loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
