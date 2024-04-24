package main

func multiplystring(s string, factor uint, ch chan<- string) {
	result := ""
	for i := uint(0); i < factor; i++ {
		result += s
	}
	ch <- result
}

func MultiplyStringsConcurrently(strings []string, factor uint) []string {
	result := []string{}
	ch := make(chan string)
	for _, s := range strings {
		go multiplystring(s, factor, ch)
		result = append(result, <-ch)
	}
	return result
}
