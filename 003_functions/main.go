package main

import "fmt"

func main() {
	// 1. Calling a simple function
	sayHello("Alice")

	// 2. Calling a function that returns a value
	sum := add(5, 3)
	fmt.Println("The sum is:", sum)

	// 3. Calling a function with multiple return values
	message, code := getStatus()
	fmt.Println("Status:", message, "Code:", code)

	// 4. Calling a function with named and naked returns
	x, y := split(17)
	fmt.Println("4. Split 17 into:", x, y)

	// 5. Function values
	// Functions can be used as arguments to other functions.
	hypot := func(x, y float64) float64 {
		return x*x + y*y
	}
	fmt.Println("5. Function Value (hypot):", hypot(5, 12))
	fmt.Println("5. Passing hypot to compute:", compute(hypot))

	// 6. Function closures
	// Each call to adder() returns a new closure bound to its own 'sum'.
	pos, neg := adder(), adder()
	fmt.Println("6. Closures (pos, neg):")
	for i := 0; i < 5; i++ {
		fmt.Printf("   %d: pos=%d, neg=%d\n", i, pos(i), neg(-2*i))
	}
}

// compute takes a function as an argument and calls it
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder returns a closure that "remembers" its own sum variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// A simple function that takes one argument and prints it
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// A function that takes two integers and returns their sum
func add(a int, b int) int {
	return a + b
}

// A function that returns multiple values
func getStatus() (string, int) {
	return "OK", 200
}

// Named return values and a "naked" return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // This is a "naked" return
}
