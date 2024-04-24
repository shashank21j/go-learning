package main

import (
	"fmt"
	"time"
)

func add(x int, y int, ch chan<- int) {
	sum := x + y
	time.Sleep(1 * time.Second)
	ch <- sum
	fmt.Println("Sum: ", sum)
}
func main() {
	// Create a channel
	ch1 := make(chan int)
	ch2 := make(chan int)
	go add(10, 20, ch1)
	go add(30, 40, ch2)

	rv1 := <-ch1
	rv2 := <-ch2

	fmt.Println("Received value1: ", rv1)
	fmt.Println("Received value2: ", rv2)
}
