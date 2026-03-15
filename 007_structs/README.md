# Structs: The Shipping Crate

In Go, a **Struct** (short for structure) is a collection of fields. It allows you to group different types of data together under one label.

---

## 1. What is a Struct?
*   **Analogy:** Imagine a **Shipping Crate** in the Amazon Warehouse. 
*   Unlike a standard variable bin (which holds one item), a struct is a large **Crate** with **internal compartments**.
*   Each compartment has a **Label** (the Field) and a **Size** (the Type).
*   For example, a `User` crate might have a compartment for a `Name` (string) and another for `Age` (int).

---

## 2. Defining a Struct
You define a struct using the `type` and `struct` keywords.

```go
type Vertex struct {
    X int
    Y int
}
```

---

## 3. Accessing Fields
You access struct fields using a **dot (`.`)**.

*   **Reading:** `fmt.Println(v.X)`
*   **Writing:** `v.X = 4`

---

## 4. Pointers to Structs
Struct fields can be accessed through a struct pointer.

*   **The Go Shortcut:** In some languages, you have to write `(*p).X` to access a field through a pointer. Go is much friendlier—you can just write **`p.X`**, and Go automatically handles the "visiting" of the address for you (it reaches into the crate for you)!

---

## 5. Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

*   `v = Vertex{1, 2}` -> Sets X to 1 and Y to 2 (order matters).
*   `v = Vertex{X: 1}` -> Sets X to 1 and Y to 0 (the zero value).
*   `v = Vertex{}` -> Sets both X and Y to 0.

---
Check out the `main.go` file in this folder to see the shipping crates in action!
