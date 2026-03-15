package main

import "fmt"

func main() {
	// Let's create a "conveyor belt" of powers of 2.
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	// 1. Getting both Index and Value
	fmt.Println("1. Full Scan (Index and Value):")
	for i, v := range pow {
		fmt.Printf("   2**%d = %d\n", i, v)
	}

	// 2. Skipping the Value
	// Use 'i := range pow' if you only need the position.
	fmt.Println("\n2. Position Only:")
	for i := range pow {
		fmt.Printf("   Scanning bin #%d...\n", i)
	}

	// 3. Skipping the Index
	// Use the blank identifier '_' to ignore the index.
	fmt.Println("\n3. Item Only:")
	sum := 0
	for _, value := range pow {
		sum += value
	}
	fmt.Printf("   Total sum of items: %d\n", sum)
}
