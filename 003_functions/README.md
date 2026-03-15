# Go Functions: The Action Takers

A **Function** is a block of code that does a specific job. Think of it as a **Mini-Machine** or a **Recipe**.

- You build the machine once.
- You can use it as many times as you want.

## 1. The "Magic" Main Function
In every Go program that you want to **run**, there is one special function that must exist: `func main()`.

*   **The Starting Gun:** This is the very first thing Go runs. 
*   **The Finish Line:** When the `main` function is finished, the entire program stops.
*   **Simple:** It never takes any inputs (arguments) and never returns any values.

## 2. What is a Function?
Imagine a **Vending Machine**.
- You put something in (Money/Arguments).
- The machine does its work (Processing).
- Something comes out (Soda/Return Value).

## 3. Inputs (Arguments)
A function can take **zero or more** arguments. 
Arguments are the pieces of information you give to the function so it can do its work. 

In the `add` example in our code, the function takes two parameters of type `int`.

### Shortening Your Arguments
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we can shorten:
`x int, y int`
to:
`x, y int`

### The "Type After Name" Rule
Notice that in Go, the **type comes after the variable name**. 
- In many other languages (like C or Java), you say `int a`. 
- In Go, we say `a int`.

### Why is the type at the end?
You might wonder why Go does this differently than Java or C. There are three simple reasons:

1.  **Readability:** It’s more natural to say "Name is a String" (`name string`) than "String Name". 
2.  **Consistency:** In Go, the name always comes first—whether it's a variable, a function, or a return type. You always know where to look for the name!
3.  **Simplicity:** In older languages, complex code could become a "spiral" that was hard to read. Go keeps everything in a straight line from left to right.

```go
func sayHello(name string) { // 'name' is the input, 'string' is the type
    fmt.Println("Hello,", name)
}
```

## 4. Outputs (Return Values)
A function can "hand back" a result to you using the `return` keyword.
*   **Analogy:** It’s like a clerk handing you a **Receipt** after you buy something.

```go
func add(a int, b int) int { // The last 'int' means this machine hands back an integer
    return a + b
}
```

## 5. Multiple Return Values (The "Go Specialty")
One of the coolest things about Go is that a function can return **more than one thing** at the same time.
*   **Analogy:** Imagine a vending machine that gives you your **Soda** AND your **Change** at the same time.

```go
func getStatus() (string, int) { // It returns a string AND an int
    return "OK", 200
}
```

## 6. Named Return Values
Go's return values may be **named**. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.
*   **Analogy:** It's like a printed receipt where every number has a label (e.g., "Tax: $1.00", "Total: $10.00").

### "Naked" Returns
A return statement without arguments returns the named return values. This is known as a **"naked" return**.

```go
func split(sum int) (x, y int) { // x and y are NAMED here
    x = sum * 4 / 9
    y = sum - x
    return // This returns x and y automatically!
}
```

> [!WARNING]
> Naked return statements should be used only in **short functions**. They can harm readability in longer functions because it becomes hard to track what is actually being returned.

## 7. Function Values
Functions are **Values** too. This is a very powerful concept!
*   **Analogy:** It’s like a **Recipe card**. You can't eat the card, but you can pass it to a friend, put it in a drawer, or hand it to a chef (the function caller).
*   Functions can be:
    1.  Passed as **Arguments** to other functions.
    2.  Returned as **Return Values** from functions.

```go
// compute takes a function as an argument!
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}
```

## 8. Function Closures
Go functions may be **Closures**. 
*   **Analogy:** Imagine a small **Robot** with its own **Internal Memory**. 
*   A closure is a function value that "remembers" and references variables from outside its body. 
*   The function is "bound" to those variables. Even if the outer function finishes, the closure keeps its own private state.

### Example: The Adder
In our code, the `adder` function returns a closure. Each time you call `adder()`, it creates a **new, separate** sum variable.

```go
func adder() func(int) int {
    sum := 0 // This is the 'Internal Memory'
    return func(x int) int {
        sum += x // The closure accesses and changes its private 'sum'
        return sum
    }
}
```

## 9. Why use Functions?
- **Reuse:** Don't write the same code twice!
- **Organization:** Break a big problem into small, easy-to-understand tasks.
- **Testing:** It's much easier to test one small "machine" than a giant factory.

---
Take a look at the `main.go` file in this folder to see these functions being called and working together!
