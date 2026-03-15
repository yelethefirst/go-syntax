# Flow Control: The Traffic Signs of Go

In programming, your code is like a **Road Trip**. Flow control statements are the **Traffic Signs**, **Detours**, and **Loops** that tell your program which way to go.

---

## 1. `for`: The Treadmill (The Only Loop)
In most languages, you have `while`, `do-while`, and `for`. Go keeps it simple: **There is only `for`.**

*   **Analogy:** A Treadmill. You keep running (looping) until you reach your goal or someone hits the stop button.
*   **The Three Parts:** 
    1.  **Init:** Start the treadmill (`i := 0`). (Optional)
    2.  **Condition:** Check if you're done (`i < 10`).
    3.  **Post:** Take a step (`i++`). (Optional)

> [!NOTE]
> **The init and post statements are optional.** If you remove them, the `for` loop behaves exactly like a `while` loop in other languages!

---

## 2. `if` and `else`: The Fork in the Road
Making decisions is the core of any program.

*   **Analogy:** You come to a fork in the road. *If* the sign says "Alice's House", you go left. *Else*, you go right.
*   **Go Feature:** `if` statements can start with a short statement to execute before the condition (like `if v := math.Pow(x, n); v < lim { ... }`). Variables declared there are only available inside the `if` block.

---

## 3. `switch`: The Remote Control
When you have many options, `if-else` chains can get messy. `switch` is much cleaner.

*   **Analogy:** A Remote Control. Each button (case) does something different. You press one, and it executes that specific command.
*   **Go Benefit:** 
    - No `break` needed! Go only runs the selected case.
    - **Switch with no condition** (same as `switch true`) is just a cleaner way to write long `if-then-else` blocks.

---

## 4. `defer`: The Sticky Note
This is one of Go's most unique and powerful features.

*   **Analogy:** You're working on a project and put a **Sticky Note** on the front door: *"Turn off the lights when you leave."* No matter how long you spend inside, you'll see the note and do the task exactly as you exit.
*   **In Practice:** `defer` is often used for "cleanup" tasks, like closing a file or disconnecting from a database.
*   **Stacking:** Multiple `defer` calls are pushed onto a stack. When the function returns, they are executed in **Last-In, First-Out (LIFO)** order.

---
Check out the `main.go` file in this folder to see these traffic signs in action!
