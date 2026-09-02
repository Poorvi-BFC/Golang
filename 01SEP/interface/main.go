package main

import "fmt"

// interface: defines a set of method rules any type with the same method can use it. here, area requires a method named area() float64 (note: spelled area, not ares)

type area interface {
	area() float64 // method — has (), takes no args, returns a string
}

// circle is a custom type with one field
// radius stores the radius value of the circle

type circle struct {
	radius int
}

// method with receiver c *circle
// this method belongs to circle and can access its data
// it is not used here because the interface method name is area

func (c *circle) area() float64 {
	return 3.14 * float64(c.radius*c.radius)
}

// main starts the program
func main() {
	// create a circle object with radius 5
	c := &circle{radius: 5}
	// call the method on the object
	c.area()
	// print the area directly using the radius value
	fmt.Println("Area of the circle is:", 3.14*float64(c.radius*c.radius))
}
