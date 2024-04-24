package main

import (
	"fmt"
)

func add(x int, y int, ch chan<- int) {
	sum := x + y
	ch <- sum
	fmt.Println("Sum: ", sum)
}
func main() {
	// Create a channel
	ch := make(chan int)
	go add(10, 20, ch)

	rv := <-ch
	fmt.Println("Received value: ", rv)

}
