//go:build variables

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := [6]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a[0]) // fetching value via index

	//slicing
	c := [5]int{1: 10, 2: 40}
	fmt.Println(c)

	//Change Elements of an Array
	b[0] = 100
	fmt.Println(b)

	//Find the Length
	fmt.Println(len(a))

}
