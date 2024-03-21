package main

import (
	"fmt"
)

func main() {
	str := " , World!"
	for _, ch := range str {
		fmt.Printf("%c %T", ch, ch)
		break
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c %T\n", str[i], str[i])
	}
	fmt.Println()
}
