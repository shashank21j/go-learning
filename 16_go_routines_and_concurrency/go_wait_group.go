package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int
	mutex sync.Mutex
}

func increment(c *counter, wg *sync.WaitGroup) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = c.value + 1
	fmt.Println("Incremented value: ", c.value)
	wg.Done()
}
func main() {
	// Go program to demonstrate the need of wait group
	wg := sync.WaitGroup{}
	wg.Add(100)
	c := counter{value: 0}

	for i := 0; i < 100; i++ {
		go increment(&c, &wg)
	}
	wg.Wait()
}
