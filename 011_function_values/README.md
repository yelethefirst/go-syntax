# Function Values: The Recipe Cards

In Go, functions are **values**. This means you can use them just like any other variable (like an `int` or a `string`).

---

## 1. What is a Function Value?
*   **Analogy:** Imagine a **Recipe Card**. 
    *   The recipe card itself isn't the meal. 
    *   But you can **pass the card** to a friend, **put it in a drawer**, or **hand it to a chef** to execute the steps.
*   In Go, you can:
    1.  Assign a function to a **variable**.
    2.  Pass a function as an **argument** to another function.
    3.  Return a function from another function.

---

## 2. Passing Functions (The Delivery Service)
Just like you can hand a recipe to a chef, you can hand a function to another function.

In our code:
*   We have a function called `compute`.
*   `compute` doesn't know *what* math it's doing yet. It just knows it needs a "recipe" (a function) that takes two numbers and returns one.
*   We can hand it `hypot` (to find the hypotenuse) or `math.Pow` (to calculate powers).

---

## 3. Why Use Function Values?
*   **Flexibility:** You can change the behavior of a program by swapping out the "recipe" without changing the main logic.
*   **Abstractions:** You can write high-level functions (like `compute`) that work with many different specific implementations.

---
Check out the `main.go` file in this folder to see the recipe cards being passed around!
