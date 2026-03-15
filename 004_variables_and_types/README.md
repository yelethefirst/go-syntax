# Variables and Types: The Amazon Warehouse

In Go, a **Variable** is like a **Storage Bin** in a massive warehouse. 
- The **Label** on the bin is the variable's **Name**.
- The **Bin Size** is determined by its **Type** (some bins are for small electronics, others are for heavy furniture).
- The **Inventory** inside is the **Value**.

## 1. Variables and Your Computer's Memory
Thinking of your computer as a giant warehouse help us understand **Memory (RAM)**:

- **Memory Address:** Every bin has a specific "Aisle and Shelf" number. You can find any piece of inventory if you know its address.
- **Static Typing:** Go is very organized. If a bin is labeled for "Books" (`string`), you can't throw a "Vespa Scooter" (`int`) into it. The system will stop you immediately!

## 2. The Garbage Collector (The Clearance Crew)
In a real warehouse, you don't want old, unsold items taking up shelf space forever. 

**How it works in practice:**
Imagine the Clearance Crew (GC) walking through the warehouse with a clipboard.
1.  **Marking (The Check):** They look at the "Boss" (your running program) to see which **Slips of Paper** (addresses) are currently being held. If a bin has a slip of paper pointing to it, the GC "Marks" it as **STAYING**.
2.  **Sweeping (The Cleanup):** Any bin that has **NO** slip of paper pointing to it is considered "unreachable". The GC "Sweeps" it.
3.  **Is it deleted?** Yes! But not necessarily by burning the item. The GC simply **removes the label** and marks the bin as **AVAILABLE**. The warehouse can now overwrite that spot with a new delivery (a new variable value).

> [!TIP]
> This is why Go is so safe—you don't have to worry about the "warehouse" getting cluttered and slowing down your computer. The Clearance Crew is always working in the background!

## 3. Scope: When do bins get cleared?
You asked: *"What happens if the program goes back to code that was already cleaned?"*

This is called **Scope**. Most variables only live as long as the function they are in.

1.  **The Function Runs:** Go assigns a bin to your variable (e.g., `age := 25`).
2.  **The Function Ends:** You leave the room. The variable is no longer needed. The Clearance Crew (GC) clears the bin.
3.  **The Function Runs AGAIN:** You come back into the room. Go doesn't try to find the "old" bin. It gives you a **Brand New Bin**, and you start over from scratch!

*   **Analogy:** Every time you call a function, it's like a **New Shift** at the warehouse. You get a fresh set of empty bins to work with. You don't have to worry about what happened during the last shift!

---

## 4. References (Important Extension)
Sometimes, a storage bin doesn't hold the item itself — it holds the **location of another bin**.

*   **Analogy:** Imagine a bin that is empty except for a **Slip of Paper**. On that paper is written: *"The TV you are looking for is in Aisle 4, Shelf 12."*
*   **Why use this?** If you have a massive item (like a whole shipping container of data), it's easier to pass around a small slip of paper with the address than to move the entire container every time!

---

## 4. Declaring Variables

There are two main ways to create variables in Go:

### Method A: Using `var`
This is the formal way. You specify the name and the type.
*   **A `var` statement can be at package or function level.** We see both in the `main.go` example in this folder.
*   **Initializers:** A `var` declaration can include initializers, one per variable (e.g., `var i, j int = 1, 2`).
*   **Type Omission:** If an initializer is present, the type can be omitted; the variable will take the type of the initializer. (e.g., `var c, python, java = true, false, "no!"`).

```go
var name string = "Alice"
var age int // This will start at 0 (the Zero Value)
var i, j = 1, 2 // Both will be 'int'
```

### Method B: Short Declaration (`:=`)
This is the most popular way to create variables **inside functions**. Go is smart enough to look at your value and guess the type.
*   **Best for:** Quick variables inside functions.

```go
country := "Nigeria" // Go knows "Nigeria" is a string!
```

---

## 5. Basic Types
Go has a rich set of built-in types:

- **`bool`**: `true` or `false`
- **`string`**: text (e.g., `"Hello"`)
- **`int`**: Whole numbers. Also specific sizes: `int8`, `int16`, `int32`, `int64`.
- **`uint`**: Unsigned (positive) whole numbers. Also: `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`.
- **`byte`**: Alias for `uint8` (good for raw data).
- **`rune`**: Alias for `int32` (represents a Unicode code point/character, e.g., '🚀').
- **`float32`, `float64`**: Decimal numbers.
- **`complex64`, `complex128`**: Complex numbers (numbers with an imaginary part).

### Factored Blocks
Just like `import` statements, you can "factor" variable declarations into blocks to group several variables together:

```go
var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

---

## 6. Zero Values (The "Empty Box" Rule)
In many languages, if you don't give a variable a value, it becomes `null` or `undefined`, which can cause crashes. 
**In Go, variables always have a value.** If you don't provide one, Go gives it a "Zero Value":

| Type | Zero Value |
| :--- | :--- |
| `int` | `0` |
| `float64` | `0.0` |
| `bool` | `false` |
| `string` | `""` (empty string) |

---

## 7. Type Conversions
The expression `T(v)` converts the value `v` to the type `T`.

In Go, **assignment between items of different type requires an explicit conversion**. You cannot simply assign an `int` to a `float64` without telling Go exactly what you are doing.

```go
var x, y int = 3, 4
var f float64 = float64(x*x + y*y)
var z uint = uint(f)
```

## 8. Type Inference
When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:
```go
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:
```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

---

## 9. Constants (`const`)
Constants are like variables, but **their value can never change** once the program starts.
*   **Analogy:** Like your **Date of Birth** or the value of **Pi**.

```go
const Gravity = 9.8
```

---

## 10. Printing Values and Types (`fmt` Verbs)
Go's `fmt` package uses "verbs" (placeholders) to format strings, similar to `printf` in other languages.

| Verb | Description | Example Output |
| :--- | :--- | :--- |
| `%v` | The **default value** in a natural format. | `(2+3i)` |
| `%T` | The **Go-syntax representation of the type**. | `complex128` |
| `%#v` | The **Go-syntax representation of the value**. | `(2+3i)` |
| `%t` | Boolean (`true/false`). | `true` |
| `%d` | Base 10 integer. | `42` |
| `%f` | Decimal point (float). | `3.141590` |
| `%p` | Pointer (memory address). | `0xc0000a2000` |
| `%s` | String. | `"Hello"` |

```go
fmt.Printf("Type: %T, Value: %v", myVar, myVar)
```

---
Check out the `main.go` file in this folder to see these in action!
