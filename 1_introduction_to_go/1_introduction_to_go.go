package main

// package declaration is mandatory
// main is the entry point of the program

import "fmt"

// import fmt package

func main() {
	fmt.Println("Hello, 世界")
	// single line comment
	/*
		multi-line comment
	*/
	// go run 1_intro.go
	// go build 1_intro.go followed by ./1_intro
	// go install 1_intro.go followed by 1_intro
}
