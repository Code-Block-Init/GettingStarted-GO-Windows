package main

import "fmt"

// assigning type - meowRect
type meowRect struct {
	// declarations
	meowLen, meowWid int
}

// function - meowArea
func (ig_a meowRect) meowArea() int {
	// calculating area of a rectangle
	return ig_a.meowLen * ig_a.meowWid
}

// function - meowPerimeter
func (ig_a meowRect) meowPerimeter() int {
	// calculating the perimeter of a rectangle
	return 2 * (ig_a.meowLen + ig_a.meowWid)
}

// main function
func main() {
	// assigning values
	my_ig := meowRect{14,5}
	// let's print
	fmt.Println("Length and Width values are: ",my_ig)
	fmt.Println("Area of the rectangle is:",my_ig.meowArea())
	fmt.Println("Perimeter of the rectangle is:",my_ig.meowPerimeter())
}
