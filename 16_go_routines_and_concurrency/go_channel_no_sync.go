package main

import (
	"fmt"
)

type counter struct {
	value int
}

func increment(c *counter, ch chan<- bool) {
	c.value = c.value + 1
	fmt.Println("Incremented value: ", c.value)
	ch <- true
}
func main() {
	// Go program to demonstrate the need of wait group
	c := counter{0}
	ch := make(chan bool)

	for i := 0; i < 100; i++ {
		go increment(&c, ch)
	}

	result := 0
	for result < 100 {
		<-ch
		result++
	}
}
