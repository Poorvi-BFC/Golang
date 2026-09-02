package main

import "fmt"

type speaker interface { //interface with speak method
	speak() string
}
type walker interface { //interface with walk method [walk()]
	walk() string
}

type Dog struct { //struct with name field
	name string
}
type pet interface { //multiple interfaces can be combined into one interface
	speaker
	walker
}

func (d Dog) speak() string { //method with receiver d Dog, where d is the instance of Dog struct. It implements the speak method of speaker interface
	return "Woof!"
}
func (d Dog) walk() string {
	return "Runs on all fours!"
}
func main() {
	dog := Dog{name: "Buddy"} //create an instance of Dog struct
	var p pet = dog           //assign the instance of Dog struct to the pet interface variable p
	fmt.Println(p.speak())
	fmt.Println(p.walk())
}
