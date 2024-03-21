package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Basic arithmatic operators
	a, b := 10, 20
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)

	// Modulus operator
	fmt.Println("a % b = ", a%b)

	// Increment and decrement operators
	a++
	fmt.Println("a++ = ", a)
	b--
	fmt.Println("b-- = ", b)

	// Assignment operators
	a += 10
	fmt.Println("a += 10 = ", a)
	b -= 10
	fmt.Println("b -= 10 = ", b)
	a *= 10
	fmt.Println("a *= 10 = ", a)
	b /= 10
	fmt.Println("b /= 10 = ", b)
	a %= 10
	fmt.Println("a %= 10 = ", a)

	// integer divided by float
	c := 10
	d := 3.0
	fmt.Println("c / d = ", float64(c)/d)

	// Power
	fmt.Println("10^2 = ", math.Pow(10, 2))

	// Min and max
	fmt.Println("Min(10, 20) = ", math.Min(10, 20))
	fmt.Println("Max(10, 20) = ", math.Max(10, 20))

	// Square root
	fmt.Println("Square root of 100 = ", math.Sqrt(100))

	// Absolute value
	fmt.Println("Absolute value of -10 = ", math.Abs(-10))

	// Round
	fmt.Println("Round 10.5 = ", math.Round(10.5))

	// Floor and Ceil
	fmt.Println("Floor 10.5 = ", math.Floor(10.5))
	fmt.Println("Ceil 10.5 = ", math.Ceil(10.5))

	// Truncate
	fmt.Println("Truncate 10.5 = ", math.Trunc(10.5))

	// Convert string to integer
	str := "1234"
	i1, _ := strconv.Atoi(str)
	fmt.Println("Convert string to integer: ", i1)

	i, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println("Convert string to integer: ", i)

	// Convert integer to string
	i2 := 1234
	str1 := strconv.Itoa(i2)
	fmt.Println("Convert integer to string: ", str1)

}
