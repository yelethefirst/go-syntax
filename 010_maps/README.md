# Maps: The Index Card System

A **Map** maps keys to values. It is Go's version of a dictionary or hash table.

---

## 1. What is a Map?
*   **Analogy:** Imagine an **Index Card System** in our warehouse.
*   Each card has a **Unique Label** (the **Key**) and some **Information** written on it (the **Value**).
*   If you know the label, you can find the card instantly. You don't have to search through a whole row of bins (like in a slice).

---

## 2. Initialization

### Using `make`
The `make` function returns a map of the given type, initialized and ready for use.
```go
m = make(map[string]int)
m["Answer"] = 42
```

### Map Literals
Map literals are like struct literals, but the keys are required.
```go
var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
```

---

## 3. Mutating Maps (Updating the Cards)

### Insert or Update
`m[key] = elem`

### Retrieve
`elem = m[key]`

### Delete
`delete(m, key)`

---

## 4. The "ok" Idiom (Checking for Cards)
If you try to get a value for a key that **doesn't exist**, Go returns the **zero value** for that type (e.g., `0` for an `int`).

This can be confusing! Did the card say `0`, or was the card missing?
To be sure, use the "two-value assignment":

```go
elem, ok := m[key]
```

*   If `key` is in `m`, `ok` is `true`.
*   If `key` is not in `m`, `ok` is `false` and `elem` is the zero value.

---

## 5. Security Note: The `nil` Map
The zero value of a map is `nil`. 

*   **Read-only:** A `nil` map has no keys, and cannot be added to.
*   **Danger:** If you try to **write** to a `nil` map (`m["key"] = 1`), your program will **panic** (crash). Always use `make` or a literal to start your map!

---
Check out the `main.go` file in this folder to see the index cards in action!
