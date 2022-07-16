package main

import "fmt"

func main() {

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println("Square area: ", squareArea)

	x:= 10
	y:= 50

	// Addition
	result:= x + y
	fmt.Println("Addition: ", result)

	// Subtraction
	result = y - x
	fmt.Println("Subtraction: ", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplication: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Modulus
	result = y % x
	fmt.Println("Modulus: ", result)

	// Increment
	x++
	fmt.Println("Increment: ", x)

	// Decrement
	x--
	fmt.Println("Decrement: ", x)

	// Reactangle area
	const rectangleBase = 10
	const rectangleHeight = 5
	rectangleArea := rectangleBase * rectangleHeight

	fmt.Println("Rectangle area: ", rectangleArea)

	// Trapeze area
	const trapezeBase = 10
	const trapezeHeight = 5
	trapezeArea := (trapezeBase + trapezeHeight) / 2

	fmt.Println("Trapeze area: ", trapezeArea)

	// Circle area
	const circleRadius = 10
	const PI = 3.14
	circleArea := PI * circleRadius * circleRadius

	fmt.Println("Circle area: ", circleArea)
}