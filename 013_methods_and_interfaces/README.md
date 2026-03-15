# Methods & Interfaces: The Deep Dive

This section covers the most important part of Go's type system: how we give types behavior and how we create flexible code using interfaces.

---

## 1. Methods are Functions
A **Method** is just a function with a special "receiver" argument.
*   **Analogy:** If a function is a **Generic Tool** (like a hammer), a method is a **Built-in Ability** of an object (like a robot having a `SwingHammer()` arm).

```go
func (v Vertex) Abs() float64 { ... }
```

---

## 2. Pointer Receivers
Why use a pointer (`*Type`) instead of a value (`Type`)?
1.  **To Modify State:** If you use a value receiver, you're working on a **copy**. Changes won't stick.
2.  **Efficiency:** To avoid copying a large struct on every call.

---

## 3. Pointer Indirection (Go Magic)
Go is friendly! 
*   If you have a **Pointer Method**, you can call it on a **Value** (`v.Scale()`). Go automatically turns it into `(&v).Scale()`.
*   If you have a **Value Method**, you can call it on a **Pointer** (`p.Abs()`). Go automatically turns it into `(*p).Abs()`.

---

## 4. Stringers (`fmt.Stringer`)
One of the most common interfaces in Go.
*   **Analogy:** An **ID Badge**. 
*   If your type has a `String() string` method, the `fmt` package will use it whenever you print the object. It's like telling Go: "Here is how I want to be introduced."

---

## 5. Errors (`error`)
Go handles errors as **Values**.
*   The `error` type is an interface. Anything that has an `Error() string` method can be treated as an error.
*   **Analogy:** A **Red Alert**. You don't "throw" it into the void; you hand it back to the caller to deal with.

---

## 6. Exercise: The IP Address
In `main.go`, you'll find a type `IPAddr [4]byte`.
1.  Try to print it without the `String()` method. It looks like a bunch of numbers in brackets.
2.  **Task:** Implement the `Stringer` interface so it prints as a "dotted quad" (e.g., `127.0.0.1`).
3.  **Answer:** Reveal the commented-out section in `main.go` to see the solution!

---
Check out the `main.go` file to see the "USB Ports" and "ID Badges" of Go in action!
