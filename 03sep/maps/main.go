package main

import "fmt"

func main() {
	var mp map[string]int
	mp["apples"] = 5
	mp["oranges"] = 10

	fmt.Println(mp)
}

// import "fmt"

// var a = make(map[string]int, 5) //empty map with capacity 5
// a["one"] = 1
// a["two"] = 2

// func main() {
// 	// fmt.Println("without any value", a)
// 	// // a := append(a) // append values to the slice.
// 	// fmt.Println("with values", a)
// 	fmt.Println(a)
// }
