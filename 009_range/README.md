# Range: The Conveyor Belt Scanner

The `range` form of the `for` loop iterates over a slice or map.

---

## 1. What is Range?
*   **Analogy:** Imagine a **Conveyor Belt Scanner** at the warehouse.
*   As each bin (item) passes by, the scanner gives you two pieces of information:
    1.  **The Index**: The "Bin Number" or position on the belt.
    2.  **The Value**: A **copy** of the item inside that bin.

---

## 2. Using Range
When you range over a slice, it returns two values for each iteration.

```go
for i, v := range mySlice {
    fmt.Printf("Bin #%d contains %v\n", i, v)
}
```

---

## 3. Skipping with the Blank Identifier (`_`)
Sometimes, you don't need both the index and the value. In Go, you can't declare a variable and not use it, so we use the **underscore** (`_`) to tell Go "I don't care about this part."

### Skipping the Value
If you only want the index:
```go
for i := range mySlice {
    fmt.Println("Processing bin", i)
}
```

### Skipping the Index
If you only want the value:
```go
for _, v := range mySlice {
    fmt.Println("Got item:", v)
}
```

---
Check out the `main.go` file in this folder to see the scanner in action!
