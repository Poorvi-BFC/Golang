package main

import "fmt"

func main() { // similar to C-lang TWIST: can handle multiple cases in a single case statement. [like OR operator]
	a := 9

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2, 3, 4, 5:
		fmt.Println("a is 2, 3, 4 or 5")
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("a is neither 1 nor 10")
	}
}
