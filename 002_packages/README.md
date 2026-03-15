# Go Packages: Organizing Your Tools

In Go, every file belongs to a **Package**. Think of a package as a **labeled box** or an **aisle in a grocery store**. 

- If you need fruit, you go to the "Fruit" aisle.
- If you need math tools, you look in the `math` package.
- If you need printing tools, you look in the `fmt` package.

## 1. The `package main` Rule
Every Go program that you want to **run** must have a file that starts with `package main`. 
*   **Analogy:** This is like the "Front Door" of your house. It tells Go exactly where to start the program.

## 2. The `import` Statement
To use tools from another package, you must "invite" them into your file using the `import` keyword.
*   **Analogy:** It’s like calling a plumber to your house. You can't use their tools until they arrive!

```go
import "fmt"  // Inviting the "fmt" (printing) tool
```

## 3. Common Standard Packages
Go comes with many built-in tools (The Standard Library). Here are the most common ones for beginners:
- **`fmt`**: Used for "Formatting" text and printing it to the screen.
- **`math`**: Used for math operations (like square roots).
- **`time`**: Used for seeing the current time or date.

## 4. The "Capital Letter" Rule (Visibility)
This is a very important rule in Go. 
- If a tool's name starts with a **CAPITAL letter** (like `Println`), it is **Public** (Exported). Anyone can use it.
- If it starts with a **lowercase letter**, it is **Private**. It can only be used inside its own package.
*   **Analogy:** Think of a house. The **Front Door** (Capital) is for guests, but the **Bedroom Door** (lowercase) is only for the people who live there.

### Examples of Exported Names (Public)
Here are some common tools you will see in Go:
- **`fmt.Println`**: The `P` is capital, so it is exported.
- **`math.Pi`**: The `P` is capital, so you can use the value of Pi.
- **`time.Now`**: The `N` is capital, used to get the current time.

### Examples of Lowercase Names (Private)
These are hidden from other packages and only used inside their own "house":
- **`math.sqrt`**: (This doesn't exist publicly!) Go uses `math.Sqrt`.
- **`fmt.newPrinter`**: This is a tool `fmt` uses internally to set up a printer. You can't use it directly because it starts with a lowercase letter.
- **`time.digits`**: A list of numbers `time` uses internally to format dates. You can't see or change it from your code.

If you tried to use `fmt.println` (with a lowercase `p`), Go would give you an error saying it cannot find that tool!

---
Now, take a look at the `main.go` file in this folder to see these packages in action!
