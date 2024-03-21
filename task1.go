package main

import (
	"math"
)

func MatchingCharacters(string1 string, string2 string) int {
	// Write your code here.
	// Return the number of matching characters between string1 and string2.
	n := math.Min(float64(len(string1)), float64(len(string2)))
	matches := 0
	for i := 0; i < n; i++ {
		if string1[i] == string2[i] {
			matches++
		}
	}
	return matches
}
