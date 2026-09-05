package main

import "fmt"

func getLast[T any](s []T) T {
	// TODO: Check if slice is empty; return zero value of type T if so
	// TODO: Return the last element of slice s
	// if len(s) == 0 {
	var zero T
	return zero
	//}
	//return s[len(s)-1]

}

// don't touch below this line

func main() {
	ints := []int{10, 20, 30}
	fmt.Println("Last int:", getLast(ints))

	strs := []string{"apple", "banana", "cherry"}
	fmt.Println("Last string:", getLast(strs))
}
