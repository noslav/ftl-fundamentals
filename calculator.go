// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Mutiply takes two values and returns the result of multuiplying the second with
// the first.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two values and returns the result of multuiplying the second with
// the first.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func Sqrt(a float64) (float64 , error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f, (sqrt of negative number is imagniary)", a)
	}
	return math.Sqrt(a) , nil 
}
