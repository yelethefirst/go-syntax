# Exercises: Concurrency

### EXERCISE 1: The Wait
Use a `sync.WaitGroup` to spawn 3 goroutines. Each should print its ID and then notify the WaitGroup it is finished.

### EXERCISE 2: The Signal
Use a channel to send a "Done" message from a child goroutine back to the main goroutine to signal completion.

---

## How to Run the Solution
```bash
go run solution/main.go
```
