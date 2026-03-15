package main

import "fmt"

func main() {
	// 1. Getting an Address (&)
	// i is a normal variable (a bin holding 42)
	i := 42
	
	// p is a pointer (a slip of paper with i's address)
	// The type of p is *int (pointer to an int)
	p := &i

	fmt.Println("1. Value of i:", i)
	fmt.Println("1. Address of i (p):", p)

	// 2. Visiting the Address (*)
	// This is called "dereferencing".
	fmt.Println("2. Value found at p (*p):", *p)

	// 3. Changing the Value through the Pointer
	// We are going to the address written on p and 
	// changing the item in that bin to 21.
	*p = 21
	fmt.Println("3. New value of i (changed via *p):", i)

	// 4. Multiple Pointers to the same bin
	j := 2701
	p = &j         // p now points to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println("4. New value of j:", j)

	// 5. Nil Pointers
	// A pointer that hasn't been assigned yet is 'nil'.
	var nilPointer *int
	fmt.Println("5. Zero value of a pointer:", nilPointer)
	
	// We check for nil before using a pointer to avoid crashes!
	if nilPointer == nil {
		fmt.Println("5. The pointer is currently nil (empty slip).")
	}

	// Once assigned, it's no longer nil
	i = 100
	nilPointer = &i
	if nilPointer != nil {
		fmt.Println("5. The pointer is no longer nil. Value:", *nilPointer)
	}
}
