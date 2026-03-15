# Arrays and Slices: The Row of Bins

In Go, both **Arrays** and **Slices** are used to store multiple items of the same type. However, they behave very differently in practice.

---

## 1. Arrays (Fixed Row of Bins)
*   **Analogy:** A concrete **Row of Bins** bolted to the floor.
*   The length is part of its type. Once you declare `[5]int`, it always has exactly 5 bins. It cannot grow or shrink.
*   **Syntax:** `[n]T` is an array of `n` values of type `T`.

```go
var a [2]string
a[0] = "Hello"
a[1] = "World"
```

---

## 2. Slices (The Dynamic Window)
*   **Analogy:** A **Window** that looks at a row of bins.
*   Unlike arrays, slices have a dynamic size. They are much more common in Go code.
*   **Syntax:** `[]T` is a slice with elements of type `T`.
*   A slice is just a **view** into an underlying array. If you change a value in the slice, you are actually changing the value in the array!

---

## 3. Slice Literals
A slice literal is like an array literal without the length.

**This is an array literal:**
```go
[3]bool{true, true, false}
```

**And this creates the same array as above, then builds a slice that references it:**
```go
[]bool{true, true, false}
```

---

## 4. Slicing Operations
You can create a slice by selecting a range from an array or another slice.

*   `a[low : high]` --> Selects bins from index `low` to `high-1`.
*   `a[:3]` --> Bins from the start to index 2.
*   `a[1:]` --> Bins from index 1 to the end.

---

## 4. Nil Slices (The "No Window" State)
The zero value of a slice is `nil`. 

*   **Analogy:** You have the **frame** of a window, but it’s not looking at any row of bins yet. 
*   A `nil` slice has:
    1.  **Length:** 0
    2.  **Capacity:** 0
    3.  **Underlying Array:** None!
*   **Safety:** You can safely call `append` on a `nil` slice. Go will see the empty window, create a new row of bins for you, and point the window there.

---

## 5. Length vs. Capacity
A slice has both a **length** and a **capacity**.

1.  **Length (`len`)**: The number of items currently in the window.
2.  **Capacity (`cap`)**: The number of items in the underlying "Row of Bins" starting from the first element in the slice.

---

## 6. Building and Growing Slices

### `make`
The `make` function creates a slice with a specific length and capacity. This is how you pre-allocate a specific number of bins.
```go
s := make([]int, 5) // len=5, cap=5
```

### `append`
The `append` function adds new items to a slice. If the underlying row of bins is too small, Go automatically creates a **bigger row** and moves all your items there!
```go
s = append(s, 10, 20)
```

---
Check out the `main.go` file in this folder to see the dynamic windows in action!
