package main

import "fmt"

//basic example of pointers in go
//func main() {
// a := 42
// b := &a
// fmt.Println("Value of a:", a)
// fmt.Println("Value of b:", b)
// *b = 21
// fmt.Println("Value of a after modifying b:", a)
//}

func getPointerValues(p *int) int {
	// TODO: Return 0 if pointer p is nil
	// TODO: Otherwise dereference p and return its int value

	if p == nil {
		return 0
	}
	return *p
}

// don't touch below this line

func main() {
	x := 42
	fmt.Println("Value of x via pointer:", getPointerValues(&x))

	var nilPtr *int
	fmt.Println("Value of nil pointer:", getPointerValues(nilPtr))
}
