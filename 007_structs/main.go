package main

import "fmt"

// 1. Defining a Struct
// A Vertex is a custom "crate" with two integer compartments: X and Y.
type Vertex struct {
	X int
	Y int
}

func main() {
	// 2. Initializing a Struct
	v := Vertex{1, 2}
	fmt.Println("1. Initial Vertex:", v)

	// 3. Accessing Fields
	// Use the dot (.) to reach inside the crate.
	v.X = 4
	fmt.Println("2. Modified X:", v.X)
	fmt.Println("2. Vertex is now:", v)

	// 4. Pointers to Structs
	// p is a slip of paper with the address of the crate 'v'.
	p := &v
	
	// Go allows us to write p.X instead of (*p).X.
	// It's like having a remote control for the crate!
	p.X = 1e9
	fmt.Println("3. Modified via pointer (p.X):", v)

	// 5. Struct Literals
	// You can specify fields by name, or leave them out to get Zero Values.
	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p1 := &Vertex{1, 2} // has type *Vertex

	fmt.Println("4. Struct Literals:")
	fmt.Printf("   v1: %v\n", v1)
	fmt.Printf("   v2: %v\n", v2)
	fmt.Printf("   v3: %v\n", v3)
	fmt.Printf("   p1 (address): %p, p1 (value): %v\n", p1, *p1)
}
