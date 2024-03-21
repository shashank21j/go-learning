package main

func main() {
	a := 10
	b := 20.01
	println("a > b = ", a > int(b))
	println("a < b = ", a < int(b))

	// complex condition
	println("a > 5 && b < 30 = ", a > 5 && b < 30)
	println("!(a > 5) || b < 30 = ", !(a > 5) || b < 30)

	c := 1
	if c == 1 {
		println("c is 1")
	} else if c == 2 {
		println("c is 2")
	} else {
		println("c is neither 1 nor 2")
	}
}
