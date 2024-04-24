package main

func KeepLongestAndShortestWord(wordSlices *[][]string) {
	// Write your code here.
	for slice := range *wordSlices {
		if len((*wordSlices)[slice]) <= 1 {
			continue
		}

		smallest := 99999
		largest := -1
		smallestIndex := -1
		largestIndex := -1
		for index, word := range (*wordSlices)[slice] {
			if len(word) < smallest {
				smallest = len(word)
				smallestIndex = index
			}
			if len(word) > largest {
				largest = len(word)
				largestIndex = index
			}
		}

		// Remove all words except the smallest and largest preserving the order
		newSlice := make([]string, 0)
		for index, word := range (*wordSlices)[slice] {
			if index == smallestIndex || index == largestIndex {
				newSlice = append(newSlice, word)
			}
		}
		(*wordSlices)[slice] = newSlice
	}
}
