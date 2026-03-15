# Concurrency: The Busy Kitchen

Go's most famous feature is how it handles multiple tasks at once. It's built for scale.

---

## 1. Goroutines: The Chefs
*   **Analogy:** Imagine a **Kitchen**. 
    *   A normal program is one chef doing everything. 
    *   A **Goroutine** is like hiring more **Chefs**. They all work in the same kitchen and share the same ingredients.
    *   Spawning a goroutine is incredibly cheap and fast.

```go
go doWork() // Run this in the background!
```

---

## 2. Channels: The Conveyor Belts
How do chefs talk to each other?
*   **Analogy:** A **Conveyor Belt**.
    *   One chef puts a dish on the belt (`c <- dish`).
    *   Another chef picks it up (`dish := <- c`).
    *   If the belt is empty, the chef **waits**. This keeps everything in sync without extra effort.

---

## 3. Select: The Multi-Tasking Waiter
*   **Analogy:** A **Waiter** standing between several conveyor belts.
*   "If Belt A has a dish, take it. If Belt B has a signal to stop, then quit." `select` lets you handle multiple asynchronous events easily.

---

## 4. Mutex: The Fridge Key
Sometimes, two chefs shouldn't touch the same thing at once (like the only Salt Shaker).
*   **Analogy:** A **Key to the Fridge**. 
*   Before you open the fridge, you `Lock()` it. When you're done, you `Unlock()`. This prevents "Race Conditions" where data gets corrupted.

---

## 5. WaitGroups: The Kitchen Bell
*   A `WaitGroup` is how we know everyone is finished. 
*   We add the number of chefs, and each chef rings a bell (`Done()`) when they finish. The main program `Wait()`s for all bells before closing the kitchen.

---

## 6. Exercise: The Web Crawler
In `main.go`, we simulate a web crawler.
*   **Task:** Can you adjust the `fetch` function to return an error if the URL is "bad"?
*   **Goal:** Learn how to pass more complex data through channels and handle synchronization with many concurrent workers.

---
Check out the `main.go` file for a full simulation of the busy Go kitchen!
