package main

import "fmt"

func doPanic() {
	panic("Panic!")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}
func main() {
	defer handlePanic()
	fmt.Println("Starting the program")
	doPanic()
	fmt.Println("Ending the program")
}
