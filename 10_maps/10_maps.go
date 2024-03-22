package main

import "fmt"

func main() {
	// Declaring a map : 1
	var my_map1 map[string]int
	fmt.Println("my_map = ", my_map1)

	// Declaring a map : 2
	my_map2 := make(map[string]int)
	fmt.Println("my_map = ", my_map2)

	// Declaring a map : 3
	my_map3 := map[string]int{}
	fmt.Println("my_map = ", my_map3)

	// Declaring a map : 4
	my_map4 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("my_map = ", my_map4)

	// Declaring a map : 5 (complex)
	my_map5 := map[string][]int{"one": {1, 2, 3}, "two": {4, 5, 6}, "three": {7, 8, 9}}
	fmt.Println("my_map = ", my_map5)

	// Declaring a map : 6 (complex)
	my_map6 := map[string]map[string]int{
		"one":   {"a": 1, "b": 2, "c": 3},
		"two":   {"d": 4, "e": 5, "f": 6},
		"three": {"g": 7, "h": 8, "i": 9},
	}
	fmt.Println("my_map = ", my_map6)
	fmt.Println(my_map6["one"]["a"])

	// Accessing a map
	elem, ok := my_map4["one"]
	fmt.Println("my_map3[\"one\"] = ", elem, ok)

	// Deleting an element from a map
	fmt.Println("my_map3 = ", my_map4)
	delete(my_map4, "one")
	fmt.Println("my_map3 = ", my_map4)

	// Iterating over a map
	for key, value := range my_map5 {
		fmt.Println(key, value)
	}

	n := 100
	divisibleMap := make(map[int]int)
	for i := 1; i <= n; i++ {
		for d := 1; d <= 5; d++ {
			if i%d == 0 {
				divisibleMap[d]++
			}
		}
	}
	fmt.Println("Divisible by 100 = ", divisibleMap)
}
