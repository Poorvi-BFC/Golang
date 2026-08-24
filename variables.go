//go:build variables

package main

import (
	"fmt"
)

// func main() {
// 	var student1 string = "John" //type is string
// 	var student2 = "Jane"        //type is inferred
// 	x := 2                       //type is inferred

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }

func main() {
	// var username string = "Jhon" // type is string; can be printed simply.
	// password := 234457892        // type is inferred

	// fmt.Println(username)
	// fmt.Println(password)

	username := "Jhon"      // to print: needs to be in string format
	password := "234457892" // type is inferred
	fmt.Println("Auth: Basic", username+":"+password)
}
