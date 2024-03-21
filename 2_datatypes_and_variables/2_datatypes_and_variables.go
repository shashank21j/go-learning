package main

import "fmt"

func main() {
	// Constants
	const y = 10
	fmt.Printf("%T %v\n", y, y)

	// Variables
	var z = 10
	fmt.Printf("%T %v\n", z, z)

	// Constant int type
	const x_int int = 10
	fmt.Printf("%T %v %d\n", x_int, x_int, x_int)

	// Constant int8 type
	const x_int8 int8 = 10
	fmt.Printf("%T %v\n", x_int8, x_int8)

	// Constant int16 type
	const x_int16 int16 = 10
	fmt.Printf("%T %v\n", x_int16, x_int16)

	// Constant int32 type
	const x_int32 int32 = 10
	fmt.Printf("%T %v\n", x_int32, x_int32)

	// Constant int64 type
	const x_int64 int64 = 10
	fmt.Printf("%T %v\n", x_int64, x_int64)

	// Constant uint type
	const x_uint uint = 10
	fmt.Printf("%T %v\n", x_uint, x_uint)

	// Constant uint8 type
	const x_uint8 uint8 = 10
	fmt.Printf("%T %v\n", x_uint8, x_uint8)

	// Constant uint16 type
	const x_uint16 uint16 = 10
	fmt.Printf("%T %v\n", x_uint16, x_uint16)

	// Constant uint32 type
	const x_uint32 uint32 = 10
	fmt.Printf("%T %v\n", x_uint32, x_uint32)

	// Constant uint64 type
	const x_uint64 uint64 = 10
	fmt.Printf("%T %v\n", x_uint64, x_uint64)

	// Constant float32 type
	const x_float32 float32 = 10.0
	fmt.Printf("%T %v\n", x_float32, x_float32)

	// Constant float64 type
	const x_float64 float64 = 10.0
	fmt.Printf("%T %v\n", x_float64, x_float64)

	// Constant complex64 type
	const x_complex64 complex64 = 10 + 10i
	fmt.Printf("%T %v\n", x_complex64, x_complex64)

	// Constant complex128 type
	const x_complex128 complex128 = 10 + 10i
	fmt.Printf("%T %v\n", x_complex128, x_complex128)

	// Constant byte type
	const x_byte byte = 10
	fmt.Printf("%T %v\n", x_byte, x_byte)

	// Constant rune type
	const x_rune rune = 10
	fmt.Printf("%T %v\n", x_rune, x_rune)

	// Constant string type
	const x_string string = "Hello, World!"
	fmt.Printf("%T %v\n", x_string, x_string)

	// Constant bool type
	const x_bool bool = true
	fmt.Printf("%T %v\n", x_bool, x_bool)
}
