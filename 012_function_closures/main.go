package main

import "fmt"

// adder returns a function that adds values to an internal sum.
// Each time adder is called, it creates a NEW, private 'sum' variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pos and neg are two SEPARATE closures.
	// They each have their own 'sum' variable trapped inside them.
	pos, neg := adder(), adder()

	fmt.Println("Counting up (pos) and down (neg):")
	for i := 0; i < 10; i++ {
		fmt.Printf(
			"Step %d: pos(%d) = %d, \t neg(%d) = %d\n",
			i, i, pos(i), 2*i, neg(-2*i),
		)
	}
}
