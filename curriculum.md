# 100 Days of Go - Curriculum

Welcome to your journey to becoming an advanced Go developer! This curriculum is divided into 5 phases, each containing 20 exercises.

---

# Phase 1 — Go Fundamentals (Exercises 1–20)

## Exercise 1 — Hello Go
Difficulty: Beginner
Goal: Write your first Go program

Description:
Install the Go compiler on your system and set up your workspace. Create a file named `main.go`. Inside this file, write a program that prints "Hello, Go!" to the console using the `fmt` package. Run it using `go run main.go`.

Concepts learned:
- Basic program structure (`package main`, `func main`)
- Importing packages
- Standard output with `fmt.Println`
- Running Go files

Challenge:
Modify the program to print your name instead of "Go!".

---

## Exercise 2 — Variables and Types
Difficulty: Beginner
Goal: Declare variables of different types

Description:
Write a program that declares variables of following types: `int`, `float64`, `string`, and `bool`. Assign them values and print them to the console along with their type using `%T` verb in `fmt.Printf`.

Concepts learned:
- Variable declaration (`var` and `:=`)
- Basic data types
- Formatted output with `fmt.Printf`
- Type inference

Challenge:
Try to assign a value of one type to a variable of another type and see what happens (e.g., assign a string to an int variable).

---

## Exercise 3 — Constants and Enumerations
Difficulty: Beginner
Goal: Use constants and iota

Description:
Create a program that defines constants for the days of the week using `iota`. Print the values of these constants. Also, define a constant for the speed of light.

Concepts learned:
- `const` keyword
- `iota` for auto-incrementing values
- Typed vs Untyped constants

Challenge:
Create a set of constants representing file sizes (KB, MB, GB, TB) using `iota` and bitwise shifts.

---

## Exercise 4 — Basic Arithmetic
Difficulty: Beginner
Goal: Perform basic mathematical operations

Description:
Write a program that takes two floating-point numbers as variables and performs addition, subtraction, multiplication, and division. Print the results clearly.

Concepts learned:
- Arithmetic operators (`+`, `-`, `*`, `/`)
- Working with `float64`
- Operator precedence

Challenge:
Calculate the area and circumference of a circle given its radius.

---

## Exercise 5 — Conditions (If/Else)
Difficulty: Beginner
Goal: Implement branching logic

Description:
Write a program that checks if a given number is even or odd. If it's even, print "Even"; otherwise, print "Odd".

Concepts learned:
- `if` and `else` statements
- Modulo operator (`%`)
- Boolean expressions

Challenge:
Modify the program to check if the number is also positive, negative, or zero.

---

## Exercise 6 — The Switch Statement
Difficulty: Beginner
Goal: Use switch for multiple conditions

Description:
Create a program that takes a number (1-7) representing a day of the week and prints the name of the day. Use a `switch` statement instead of multiple `if` statements.

Concepts learned:
- `switch` statement
- `case` and `default`
- Implicit break in Go

Challenge:
Use a `switch` statement without an expression (acting like an if-else chain) to categorize a test score (e.g., A, B, C, D, F).

---

## Exercise 7 — For Loops
Difficulty: Beginner
Goal: Iterate using the only loop in Go

Description:
Write a program that prints numbers from 1 to 10 using a `for` loop.

Concepts learned:
- Standard `for` loop structure (init; condition; post)
- Incrementing variables

Challenge:
Print only the even numbers from 1 to 20.

---

## Exercise 8 — While-style For Loop
Difficulty: Beginner
Goal: Use a for loop like a while loop

Description:
In Go, there is no `while` keyword. Use a `for` loop with only a condition to print powers of 2 less than 100.

Concepts learned:
- `for` loop with condition only
- Comparative operators

Challenge:
Implement a simple "guess the number" loop where the loop continues until a variable matches a target value.

---

## Exercise 9 — Infinite Loops and Break
Difficulty: Beginner
Goal: Control flow with break and continue

Description:
Write an infinite loop that increments a counter. Inside the loop, use an `if` statement to `break` the loop when the counter reaches 5. Use `continue` to skip printing the number 3.

Concepts learned:
- `for {}` (infinite loop)
- `break` statement
- `continue` statement

Challenge:
Create a loop that prints the first 5 prime numbers.

---

## Exercise 10 — Basic Functions
Difficulty: Beginner
Goal: Define and call simple functions

Description:
Write a function `multiply` that takes two integers as parameters and returns their product. Call this function from `main` and print the result.

Concepts learned:
- Function declaration
- Parameters and return types
- Function calls

Challenge:
Write a function that calculates the factorial of a number.

---

## Exercise 11 — Multiple Return Values
Difficulty: Beginner
Goal: Return more than one value from a function

Description:
Write a function `divide` that takes two integers. It should return the quotient and the remainder.

Concepts learned:
- Multiple return values
- Short variable declaration for multiple returns

Challenge:
Write a function that returns the minimum and maximum of three given integers.

---

## Exercise 12 — Named Return Values
Difficulty: Beginner
Goal: Improve code readability with named returns

Description:
Modify the `divide` function from Exercise 11 to use named return values.

Concepts learned:
- Named return parameters
- Naked `return` statement

Challenge:
Write a function `split` that takes an integer representing total cents and returns named values for `dollars` and `cents`.

---

## Exercise 13 — User Input (fmt.Scan)
Difficulty: Beginner
Goal: Capture input from the console

Description:
Write a program that asks the user for their name and age using `fmt.Scan`. Print a greeting that includes these details.

Concepts learned:
- Pointers (basic usage for `Scan`)
- `fmt.Scan` and `fmt.Scanln`
- Reading standard input

Challenge:
Read multiple space-separated words into different variables in one go.

---

## Exercise 14 — More Complex Input (bufio)
Difficulty: Beginner
Goal: Read full lines of input

Description:
`fmt.Scan` stops at spaces. Use the `bufio` package and `os.Stdin` to read a full line of text (including spaces) from the user.

Concepts learned:
- `bufio.NewReader`
- `os.Stdin`
- `ReadString` method

Challenge:
Create a simple program that echoes back whatever the user types until they type "exit".

---

## Exercise 15 — Type Conversion
Difficulty: Beginner
Goal: Convert between different data types

Description:
Write a program that takes a `float64` value and converts it to an `int`. Also, convert an `int` to a `float64`. Note how the values change (truncation).

Concepts learned:
- Explicit type conversion (`type(value)`)
- Precision loss during conversion

Challenge:
Read a string input (e.g., "123") and convert it to an integer using the `strconv` package.

---

## Exercise 16 — Basic String Manipulation
Difficulty: Beginner
Goal: Work with the strings package

Description:
Write a string program that prints its length, converts it to uppercase, and checks if it contains a specific substring. Use the `strings` package.

Concepts learned:
- `len()` function
- `strings` package functions (`ToUpper`, `Contains`)
- String immutability

Challenge:
Replace all occurrences of a word in a sentence with another word.

---

## Exercise 17 — The Math Package
Difficulty: Beginner
Goal: Use standard library math functions

Description:
Create a program that calculates the square root of a number, the absolute value of a negative number, and rounds a decimal to the nearest integer.

Concepts learned:
- `math` package
- `math.Sqrt`, `math.Abs`, `math.Round`

Challenge:
Calculate the hypotenuse of a right-angled triangle using the Pythagorean theorem with `math.Pow`.

---

## Exercise 18 — Scoping
Difficulty: Beginner
Goal: Understand variable scope

Description:
Write a program with a package-level variable and a function-level variable with the same name. Observe which one is accessed inside the function.

Concepts learned:
- Package scope vs. Local scope
- Variable shadowing

Challenge:
Create a nested block (`{ ... }`) inside a function and see if variables declared inside it are accessible outside.

---

## Exercise 19 — Simple Calculator
Difficulty: Beginner
Goal: Combine multiple concepts into a small program

Description:
Build a command-line calculator that asks for two numbers and an operator (+, -, *, /). Use a `switch` statement to perform the operation and print the result.

Concepts learned:
- Input handling
- Control flow
- Functions (optional: put operations in functions)

Challenge:
Add error handling to prevent division by zero.

---

## Exercise 20 — Palindrome Checker
Difficulty: Beginner
Goal: Implement a basic algorithm

Description:
Write a function that checks if a given string is a palindrome (reads the same forwards and backwards).

Concepts learned:
- String iteration
- Logic for reversing or comparing strings

Challenge:
Implement the check without creating a new reversed string (use two pointers).
