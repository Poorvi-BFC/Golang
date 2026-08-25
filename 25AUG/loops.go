//go:build variables

package main

import "fmt"

//Note: continue and break are usually used with conditions.

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		// if i == 3 {             // if & continue statement implemented
		// 	fmt.Println("haha.")
		// 	continue
		// }
		fmt.Println("*")
	}
	// // Nested for loop
	size := [3]string{"small", "mid", "big"}
	fruit := [4]string{"apple", "banana", "mango", "orange"}
	for i := 0; i < len(size); i++ {
		for j := 0; j < len(fruit); j++ {
			fmt.Println(size[i], fruit[j])
		}
	}
}

// range topic is left for later. It is used to iterate over arrays, slices, maps, strings, and channels. It returns two values: the index and the value of the element at that index.
