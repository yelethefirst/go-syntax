package main

// We "import" packages to use their tools
import (
	"fmt"  // The "Printing" package
	"math" // The "Math" package
	"time" // The "Time" package
)

func main() {
	// 1. Using the 'fmt' package to print text
	fmt.Println("Hello from the Packages lesson!")

	// 2. Using the 'math' package to calculate
	// Note how Sqrt starts with a CAPITAL letter!
	number := 16.0
	result := math.Sqrt(number)
	fmt.Printf("The square root of %.0f is %.0f\n", number, result)

	// 3. Using the 'time' package to see the clock
	// Note how Now() starts with a CAPITAL letter!
	rightNow := time.Now()
	fmt.Println("The current time is:", rightNow.Format("15:04:05"))
}
