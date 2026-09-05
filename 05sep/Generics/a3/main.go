package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

func add[T Number](a, b T) T {
	// TODO: Return the sum of a and b
	sum := a + b
	return sum
}

// don't touch below this line

func main() {
	fmt.Println("Int sum:", add(10, 20))
	fmt.Println("Float sum:", add(3.14, 2.71))
}
