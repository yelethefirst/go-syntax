package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// 1. Initializing with make
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println("1. Map after make:", m["Bell Labs"])

	// 2. Map Literals
	fmt.Println("\n2. Map Literals:")
	locations := map[string]Vertex{
		"Google": {37.42202, -122.08408},
		"Apple":  {37.33182, -122.03118},
	}
	fmt.Println("   Locations:", locations)

	// 3. Mutating Maps
	fmt.Println("\n3. Mutating Maps:")
	mm := make(map[string]int)

	mm["Answer"] = 42 // Insert
	fmt.Println("   The value:", mm["Answer"])

	mm["Answer"] = 48 // Update
	fmt.Println("   The updated value:", mm["Answer"])

	delete(mm, "Answer") // Delete
	fmt.Println("   The value after delete:", mm["Answer"])

	// 4. Checking for Existence (The ok Idiom)
	fmt.Println("\n4. The 'ok' Idiom:")
	v, ok := mm["Answer"]
	fmt.Printf("   Value: %v, Present? %v\n", v, ok)

	mm["Answer"] = 100
	v, ok = mm["Answer"]
	fmt.Printf("   Value: %v, Present? %v\n", v, ok)

	// 5. Printing a map (using %v, %T)
	fmt.Println("\n5. Formatting Showcase:")
	fmt.Printf("   Type: %T\n", locations)
	fmt.Printf("   Value (%%v): %v\n", locations)
	fmt.Printf("   Value (%%#v): %#v\n", locations)
}
