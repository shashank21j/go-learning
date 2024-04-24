package main

import (
	"fmt"
	"time"
)

func add(x int, y int, delay int, ch chan<- int) {
	sum := x + y
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- sum
	fmt.Println("Sum: ", sum)
}
func main() {
	// Create a channel
	ch1 := make(chan int)
	ch2 := make(chan int)
	go add(10, 20, 1, ch1)
	go add(30, 40, 1, ch2)

	select {
	case rv1 := <-ch1:
		fmt.Println("Received value1: ", rv1)
	case rv2 := <-ch2:
		fmt.Println("Received value2: ", rv2)
	}
}
