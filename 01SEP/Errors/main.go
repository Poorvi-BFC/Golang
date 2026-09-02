package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := Divide(10, 2)
	// ALWAYS check err right after the call — this is the Go convention
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result) // 5

	// result, err = Divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err) // Error: cannot divide by zero
	// 	return
	// }
	// fmt.Println("Result:", result) // never reached
}
