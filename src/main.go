package main

import "fmt"

func main() {
	// Init constants
	const hello = "Hello, World!"
	const pi float64 = 3.141592653589793
	const pi2 = 3.141592653589793

	// Init integer variables
	base := 17
	var height int = 14
	var area int

	fmt.Println("Base: ", base)
	fmt.Println("Height: ", height)
	fmt.Println("Area: ", area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println("Square area: ", squareArea)

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)
}