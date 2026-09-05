package main

import "fmt"

func getDailyUserLimits(signup []int) []int {
	// TODO: Preallocate a slice using make() with the same length as signup
	limits := make([]int, len(signup))
	// TODO: Loop over signup and set each element to double its value
	for i, v := range signup {
		limits[i] = v * 2
	}
	// TODO: Return the preallocated slice
	return limits
}

func main() {
	signup := []int{5, 12, 8, 20}
	limits := getDailyUserLimits(signup)

	fmt.Println("Signup:", signup)
	fmt.Println("Limits: ", limits)
}
