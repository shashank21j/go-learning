package main

import (
	"fmt"
)

func main() {
	// Create a channel
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	fmt.Println("Received value1: ", <-ch)
	fmt.Println("Received value2: ", <-ch)

}
