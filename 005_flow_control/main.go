package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// 1. The Standard 'for' Loop
	// The treadmill analogy: run 10 times.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("1. For loop sum:", sum)

	// 2. 'for' as a 'while' loop
	// The init and post statements are optional.
	// This makes 'for' behave exactly like 'while' in other languages.
	count := 1
	for count < 100 {
		count += count
	}
	fmt.Println("2. For (without init/post):", count)

	// 3. 'if' with a Short Statement
	// Like variables in 'for' loops, variables declared in 'if'
	// are scoped to the 'if' (and 'else') block.
	if v := math.Pow(3, 2); v < 10 {
		fmt.Printf("3. Short If: %v is less than 10\n", v)
	} else {
		fmt.Printf("3. Short If: %v is NOT less than 10\n", v)
	}

	// 4. 'switch' - The Remote Control
	// No 'break' needed in Go!
	fmt.Print("4. Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// 5. 'switch' with no condition
	// Switch without a condition is the same as 'switch true'.
	// This is a cleaner way to write long if-then-else chains.
	t := time.Now()
	switch { // same as 'switch true'
	case t.Hour() < 12:
		fmt.Println("5. Good morning!")
	case t.Hour() < 17:
		fmt.Println("5. Good afternoon!")
	default:
		fmt.Println("5. Good evening!")
	}

	// 6. 'defer' - The Sticky Note
	// Defer statement defers the execution of a function until the
	// surrounding function returns.
	defer fmt.Println("6. Defer: I ran at the VERY end!")
	fmt.Println("6. Defer: I was called first, but I'll print before the deferred message.")

	// 7. Stacking 'defer'
	// Deferred function calls are pushed onto a stack (LIFO).
	fmt.Println("7. Stacking Defer: Counting...")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("7. Defer Stack: %d\n", i)
	}
	fmt.Println("7. Stacking Defer: Done (wait for returns...)")
}
