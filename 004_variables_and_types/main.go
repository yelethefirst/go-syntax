package main

import (
	"fmt"
	"math/cmplx"
)

// 0. Factored Var Blocks
// Variable declarations may be "factored" into blocks,
// just like import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 0. Package Level Declaration
// 'var' can be used here, at the package level.
// Short declaration (:=) CANNOT be used here.
var packageLevelVar = "I am at the top!"

func main() {
	// A 'var' statement can be at package or function level.
	// We see both in this example.

	// 1. Function Level Declaration
	var functionLevelVar = "I am inside a function!"
	fmt.Println("0. Package Level:", packageLevelVar)
	fmt.Println("1. Function Level:", functionLevelVar)

	// 2. Declaring with 'var' (with explicit types)
	var name string = "Alice"
	var age int = 25
	fmt.Printf("2. Using 'var': Name is %s, Age is %d\n", name, age)

	// 3. Variables with Initializers
	// A 'var' declaration can include initializers, one per variable.
	var i, j int = 1, 2
	fmt.Println("3. Initializers:", i, j)

	// 4. Type Omission
	// If an initializer is present, the type can be omitted.
	// Go will infer the type from the value.
	var c, python, java = true, false, "no!"
	fmt.Println("4. Type Omission:", c, python, java)

	// 5. Short Variable Declaration (:=)
	// This is the most common way to declare variables INSIDE functions.
	// Go "guesses" (infers) the type for you.
	country := "Nigeria" // Go knows this is a string
	isStudent := true     // Go knows this is a bool
	fmt.Printf("2. Using ':=': Country is %s, Student: %t\n", country, isStudent)

	// 3. Zero Values
	// In Go, variables are never "undefined" or "null" by default. 
	// If you don't give them a value, Go gives them a default "Zero Value".
	var emptyString string
	var emptyInt int
	var emptyBool bool
	fmt.Printf("3. Zero Values: String: [%s], Int: %d, Bool: %t\n", emptyString, emptyInt, emptyBool)

	// 9. Basic Types Showcase
	// This shows the variables from the factored block above
	fmt.Printf("9. Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("9. Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("9. Type: %T Value: %v\n", z, z)

	// Local basic types
	var b byte = 'A'  // alias for uint8
	var r rune = '🚀' // alias for int32 (Unicode code point)
	fmt.Printf("9. Byte: %v (%c), Rune: %U (%c)\n", b, b, r, r)

	// 10. Type Conversions
	// The expression T(v) converts the value v to the type T.
	// Unlike some other languages, Go requires EXPLICIT conversion.
	var x, y int = 3, 4
	var f float64 = float64(x*x + y*y)
	var z_conv uint = uint(f)
	fmt.Printf("10. Conversions: x=%v, y=%v, f=%v, z=%v\n", x, y, f, z_conv)

	// 11. Type Inference
	// When declaring a variable without an explicit type (using := or var =),
	// the variable's type is inferred from the value on the right hand side.
	v := 42           // int
	w := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("11. Inference: v is %T, w is %T, g is %T\n", v, w, g)

	// 12. Constants
	// Constants are values that NEVER change.
	const Gravity = 9.8
	fmt.Println("12. Constant Gravity:", Gravity)

	// 13. Formatting Showcase (fmt verbs)
	// Go uses "verbs" to format strings.
	fmt.Printf("\n13. Formatting Examples:\n")
	fmt.Printf("Default value (%%v): %v\n", z)
	fmt.Printf("Go syntax representation (%%#v): %#v\n", z)
	fmt.Printf("Type of the value (%%T): %T\n", z)
	fmt.Printf("Pointer address (%%p): %p\n", &z)
	fmt.Printf("Percent sign (%%%%): 100%%\n")
}
