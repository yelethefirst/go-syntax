package main

import (
	"fmt"
	"sync" // The "Synchronization" package
	"time"
)

/*
THE SYNC LIBRARY: The Kitchen Safety Manual

While Go's philosophy is "Don't communicate by sharing memory; share memory by communicating" (using Channels), 
the 'sync' package provides low-level tools for when you HAVE to share memory safely.

Common Tools in the sync library:
1. sync.WaitGroup: The "Kitchen Bell". Use this to wait for a collection of goroutines to finish.
   - .Add(n): "I'm hiring 'n' more chefs."
   - .Done(): "I'm finished with my task!" (Rings the bell).
   - .Wait(): "Don't close the kitchen until everyone has rung their bell."

2. sync.Mutex: The "Fridge Key" (Mutual Exclusion). Use this to ensure only one goroutine touches a variable at a time.
   - .Lock(): "I have the key! Nobody else can touch the fridge."
   - .Unlock(): "I'm done. Here's the key for the next person."

3. sync.Once: The "One-Time Setup". Ensures a piece of code (like initializing a database) runs exactly once, 
   no matter how many goroutines try to call it.

When to use 'sync' vs 'channels'?
- Use CHANNELS when you are passing data around (The Conveyor Belt).
- Use SYNC when you are protecting a shared variable (The Shared Fridge).
*/

// --- 1. Goroutines ---
// A goroutine is a lightweight thread managed by the Go runtime.
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup that this chef is finished
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// --- 2. Channels ---
// Channels are the pipes that connect concurrent goroutines.
// You can send values from one goroutine into channels, and receive those values into another.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

// --- 3. Close and Range ---
// 'close' is a BUILT-IN function (no library needed).
// A sender can close a channel to indicate that no more values will be sent.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // Always close from the sender's side
}

// --- 4. Select ---
// Select lets a goroutine wait on multiple communication operations.
func waiter(c, quit chan int) {
	for {
		select {
		case x := <-c:
			fmt.Println("Received:", x)
		case <-quit:
			fmt.Println("Quit signal received. Closing kitchen.")
			return
		}
	}
}

// --- 5. Mutex ---
// Mutex (mutual exclusion) ensures only one goroutine can access a variable at a time.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) string {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return fmt.Sprintf("%d", c.v[key])
}

// --- EXERCISE: Simple Crawler ---
// Use everything you've learned to simulate a crawler that fetches data concurrently.
type Result struct {
	url  string
	body string
}

func fetch(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate web fetch
	time.Sleep(200 * time.Millisecond)
	results <- Result{url, "Body content for " + url}
}

func main() {
	fmt.Println("--- 1. Goroutines & WaitGroup ---")
	// The Busy Kitchen: Spawning Chefs
	var wg sync.WaitGroup
	wg.Add(2)
	go say("Chef 1: Chopping Veggies", &wg)
	go say("Chef 2: Stirring Soup", &wg)
	wg.Wait()

	fmt.Println("\n--- 2. Channels ---")
	// The Conveyor Belt
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println("Sum parts received:", x, y, "Total:", x+y)

	fmt.Println("\n--- 3. Fibonacci with Range/Close ---")
	ch := make(chan int, 10) // Buffered channel
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Println("\n--- 4. Select ---")
	orderChan := make(chan int)
	quitChan := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			orderChan <- i
		}
		quitChan <- 0
	}()
	waiter(orderChan, quitChan)

	fmt.Println("\n--- 5. Mutex (Safe Counter) ---")
	counter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.Inc("clicks")
	}
	time.Sleep(time.Second) // Wait for goroutines
	fmt.Println("Counter Value:", counter.Value("clicks"))

	fmt.Println("\n--- 6. Exercise: Concurrent Crawler ---")
	urls := []string{"google.com", "golang.org", "github.com"}
	results := make(chan Result, len(urls))
	var crawlerWg sync.WaitGroup

	for _, url := range urls {
		crawlerWg.Add(1)
		go fetch(url, results, &crawlerWg)
	}

	crawlerWg.Wait()
	close(results)

	for res := range results {
		fmt.Printf("Fetched %s: %s\n", res.url, res.body)
	}
}
