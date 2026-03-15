package main

import (
	"fmt"
	"math"
	"time"
)

// Vertex is our core type for demonstration.
type Vertex struct {
	X, Y float64
}

// --- 1. Methods are Functions ---
// A method is just a function with a special receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This is the function equivalent of the method above.
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// --- 2. Pointer Receivers vs. Value Receivers ---
// Scale uses a pointer receiver to modify the original vertex.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc is the function equivalent. It MUST take a pointer.
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// --- 3. Pointer Indirection ---
// Go automatically handles the conversion between values and pointers for methods.
// This is called 'Indirection'.

// --- 4. Stringers ---
// The fmt package looks for the Stringer interface to decide how to print a type.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// --- 5. Errors ---
// The error type is an interface similar to fmt.Stringer.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// --- EXERCISE: IP Address Stringer ---
// Task: Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// e.g., IPAddr{127, 0, 0, 1} should print as "127.0.0.1".

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// --- EXERCISE ANSWER (Implementation) ---
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	v := Vertex{3, 4}

	fmt.Println("--- 1. Methods vs Functions ---")
	fmt.Printf("Method: %v, Function: %v\n", v.Abs(), AbsFunc(v))

	fmt.Println("\n--- 2. Pointer Receivers ---")
	v.Scale(10)
	fmt.Printf("After Scale Method: %v\n", v)
	ScaleFunc(&v, 2)
	fmt.Printf("After Scale Function: %v\n", v)

	fmt.Println("\n--- 3. Pointer Indirection ---")
	// Methods with pointer receivers take either a value or a pointer as the receiver when they are called.
	p := &v
	p.Scale(0.5) // OK
	v.Scale(0.1) // OK: Go interprets v.Scale(0.1) as (&v).Scale(0.1)
	fmt.Printf("After Indirection (Value calling Pointer Method): %v\n", v)

	fmt.Println("\n--- 4. Stringers ---")
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	fmt.Println("\n--- 5. Errors ---")
	if err := run(); err != nil {
		fmt.Println("Caught Error:", err)
	}

	fmt.Println("\n--- 6. Exercise: IPAddr Stringer ---")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
