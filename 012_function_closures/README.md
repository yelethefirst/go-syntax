# Function Closures: The Robot with a Secret Journal

A **Closure** is a function value that "remembers" variables from outside its body.

---

## 1. What is a Closure?
*   **Analogy:** Imagine a small **Robot**. 
    *   This robot has its own **Secret Journal** (internal memory).
    *   When you create the robot, it starts with a clean journal.
    *   Every time you give it a command, it can read and write to its journal.
    *   If you build a second robot, it gets its **own separate journal**. What Robot A writes doesn't affect Robot B.

In Go:
The function is "bound" to variables outside its scope. Even after the original function (`adder`) finishes, the closure keeps those variables alive.

---

## 2. In Our Code: The Adder
*   `adder()` defines a variable `sum := 0`.
*   It returns a function that increases that `sum`.
*   When we call `pos := adder()`, the `pos` function "closes over" that specific `sum` variable.
*   It is now the only one who can change it.

---

## 3. Why Use Closures?
*   **Data Isolation:** It’s a great way to give a function "private state" without using global variables.
*   **Consistency:** Each instance of the closure starts fresh and maintains its own history.

---
Check out the `main.go` file in this folder to see the robots updating their secret journals!
