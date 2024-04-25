package main

import "fmt"

// Named return values
func test(x int, y int) (a int, b int) {
	a = x + y
	b = x - y
	return
}

// Variadic function
func sumIntegers(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Func as return value
func getFunc(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Func as argument
func callFunc(callable func(string) string, arg string) {
	result := callable(arg)
	fmt.Println("Result: ", result)
}

func main() {
	a, b := test(10, 20)
	fmt.Println("a + b = ", a)
	fmt.Println("a - b = ", b)

	fmt.Println("Sum of 1, 2, 3, 4, 5 = ", sumIntegers(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of nums = ", sumIntegers(nums...))

	add10 := getFunc(10)
	fmt.Println("add10(20) = ", add10(20))

	pos := adder()
	neg := adder()
	for i := 0; i < 10; i++ {
		fmt.Println("pos(i) = ", pos(i), "neg(-i) = ", neg(-2*i))
	}
	myFunc := func(s string) string {
		return "Hello, " + s
	}
	callFunc(myFunc, "World")
}
