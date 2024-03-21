package main

import "fmt"

func DailyRoutine(time string) string {
	// Write your code here.
	switch time {
	case "morning":
		return "Time to get up!"
	case "afternoon":
		return "Time for lunch!"
	case "night":
		return "Time to sleep!"
	default:
		return "Have a great day :)"
	}
}

func main() {
	// Write your code here.
	fmt.Println(DailyRoutine("morning"))
	fmt.Println(DailyRoutine("afternoon"))
	fmt.Println(DailyRoutine("night"))
	fmt.Println(DailyRoutine("midnight"))
}
