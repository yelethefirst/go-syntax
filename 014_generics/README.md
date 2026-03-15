# Generics: The Shape Sorter

Before Generics, if you wanted a function to work with `int` and `string`, you often had to write it twice. Now, you can write it once!

---

## 1. What are Generics?
*   **Analogy:** Imagine a **Shape Sorter** toy.
    *   Without generics, you have a box that *only* fits squares. 
    *   With generics, you have a **Magic Box** where you can say: "This box will fit whatever shape I decide to put in right now."
*   In Go, we use **Type Parameters** inside square brackets `[T any]`.

---

## 2. Generic Functions
A function can take a type as a "hidden argument".
*   `Index[T comparable]` means: "This function works for any type `T`, as long as `T` can be compared using `==`."

---

## 3. Generic Types
You can also make structs generic.
*   **Analogy:** A **Lunchbox**. It can hold a sandwich, or it can hold an apple. The lunchbox is the same, but the **Content Type** changes.
*   In our code, we use a `List[T]` which is a linked list that can hold any type of data.

---

## 4. Why Use Generics?
*   **Dry (Don't Repeat Yourself):** No more copy-pasting code for different types.
*   **Type Safety:** Unlike using `interface{}`, the compiler knows exactly what type is inside your `List[T]` at compile time.

---
Check out the `main.go` file to see the Magic Box in action!
