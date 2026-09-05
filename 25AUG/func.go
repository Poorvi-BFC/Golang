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
	// myName()                     // function calling.
	// myFunc("Alice")              // function calling with argument.
	// result := returnFunc(10, 11) // return function calling with argument.
	//fmt.Println("Result:", result)
	fmt.Println(multipleReturnFunc()) // multiple return function calling.
}

//Note: Functions can return values. The return type is specified after the parameter list. For example, func add(a int, b int) int { return a + b } is a function that takes two integers as parameters and returns an integer.

func returnFunc(a int, b int) (result int) {
	result = a + b
	return result
}

// Gaurd clause is a programming pattern that is used to simplify the code by handling edge cases or exceptional conditions at the beginning of a function, rather than nesting the main logic of the function inside multiple if statements. This can make the code easier to read and understand, as it reduces the amount of indentation and makes it clear what conditions must be met for the main logic to execute.

// go can return multiple return values from a function. The return type is specified after the parameter list.

func multipleReturnFunc() (string, string, string) { // multiple return function creation.
	return "golang", "java", "python" // multiple return values
}
