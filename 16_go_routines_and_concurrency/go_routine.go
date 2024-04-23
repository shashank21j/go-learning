package main

import (
	"fmt"
	"time"
)

func run1() {
	time.Sleep(1 * time.Second)
	fmt.Println("run1")
}

func run2() {
	time.Sleep(2 * time.Second)
	fmt.Println("run2")
}

func run3() {
	time.Sleep(3 * time.Second)
	fmt.Println("run3")
}

func main() {
	go run1()
	go run2()
	go run3()
	time.Sleep(4 * time.Second)
	fmt.Println("done")
}
