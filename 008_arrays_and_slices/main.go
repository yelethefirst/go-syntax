package main

import "fmt"

func main() {
	// 1. Arrays (Fixed Size)
	// a is a row of 2 strings. Bolted to the floor!
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("1. Array a:", a[0], a[1])
	fmt.Println("1. Full array:", a)

	// 2. Slices (The Window)
	// names is an array of 4 items.
	names := [4]string{
		"Alice",
		"Bob",
		"Charlie",
		"Diana",
	}
	fmt.Println("\n2. Original names array:", names)

	// Create two "windows" (slices) looking at the same array.
	x := names[0:2] // Alice, Bob
	y := names[1:3] // Bob, Charlie
	fmt.Println("2. Slice x (0:2):", x)
	fmt.Println("2. Slice y (1:3):", y)

	// Modification through a slice affects the underlying array!
	y[0] = "Robert"
	fmt.Println("2. Modified y[0] to 'Robert'")
	fmt.Println("2. New x:", x)          // Bob has changed to Robert
	fmt.Println("2. New array:", names) // The original array is changed

	// 3. Slice Literals
	// A slice literal is like an array literal without the length.
	
	// This is an array literal (fixed size 3)
	arrayLiteral := [3]bool{true, true, false}
	
	// This creates the same array, then builds a slice that references it:
	sliceLiteral := []bool{true, true, false}
	
	fmt.Println("\n3. Array Literal:", arrayLiteral)
	fmt.Println("3. Slice Literal:", sliceLiteral)

	// Another example with integers
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("\n3. Slice Literal q:", q)

	// 4. Length and Capacity
	printSlice("4. Initial q", q)

	// Slice q to give it zero length.
	q = q[:0]
	printSlice("4. Empty q", q)

	// Extend its length.
	q = q[:4]
	printSlice("4. Extended q", q)

	// Drop its first two values.
	q = q[2:]
	printSlice("4. Dropped first two", q)

	// 5. Using make
	s := make([]int, 5) // len=5, cap=5
	printSlice("5. Using make(5)", s)

	// 6. Nil Slices
	// A slice that hasn't been initialized is 'nil'.
	var nilSlice []int
	printSlice("6. Nil slice", nilSlice)
	if nilSlice == nil {
		fmt.Println("6. nilSlice is truly nil!")
	}

	// 7. Growing with append
	var numbers []int
	printSlice("7. Starting with nil", numbers)

	// append works on nil slices!
	numbers = append(numbers, 0)
	printSlice("7. Appended 0", numbers)

	// The slice grows as needed.
	numbers = append(numbers, 1)
	printSlice("7. Appended 1", numbers)

	// We can add multiple elements at once.
	numbers = append(numbers, 2, 3, 4)
	printSlice("7. Appended 2,3,4", numbers)
}

func printSlice(label string, s []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", label, len(s), cap(s), s)
}
