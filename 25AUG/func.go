package main

import "fmt"

func myName() { // function creation.
	fmt.Println("My name is Kio")
}
func myFunc(fname string) { //func FunctionName(param1 type, param2 type){} // function creation with parameter.
	fmt.Println("My name is", fname)
}

//When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while "Alice" is an argument.

func main() {
	myName()                     // function calling.
	myFunc("Alice")              // function calling with argument.
	result := returnFunc(10, 11) // return function calling with argument.
	fmt.Println("Result:", result)
}

//Note: Functions can return values. The return type is specified after the parameter list. For example, func add(a int, b int) int { return a + b } is a function that takes two integers as parameters and returns an integer.

func returnFunc(a int, b int) (result int) {
	result = a + b
	return result
}
