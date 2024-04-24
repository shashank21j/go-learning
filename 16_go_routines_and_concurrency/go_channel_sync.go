package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int
	mutex sync.Mutex
}

func increment(c *counter, ch chan<- bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = c.value + 1
	fmt.Println("Incremented value: ", c.value)
	ch <- true
}
func main() {
	// Go program to demonstrate the need of wait group
	c := counter{value: 0}
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
