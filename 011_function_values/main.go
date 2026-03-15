package main

import (
	"fmt"
	"math"
)

// compute is a function that takes another function as an argument.
// This shows that functions are values that can be passed around!
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 1. Assigning a function to a variable
	// hypot is now a variable that holds a function value.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("Result of hypot(5, 12):", hypot(5, 12))

	// 2. Passing a function value as an argument
	// We pass our 'hypot' function to 'compute'.
	fmt.Println("Result of compute(hypot):", compute(hypot))

	// 3. Passing a built-in function as an argument
	// math.Pow is also a function value!
	fmt.Println("Result of compute(math.Pow):", compute(math.Pow))
}
