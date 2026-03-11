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

# Phase 2 — Intermediate Go (Exercises 21–40)

## Exercise 21 — Arrays
Difficulty: Intermediate
Goal: Understand fixed-size sequences

Description:
Declare an array of 5 integers and initialize it with values. Write a loop to calculate the sum of all elements in the array.

Concepts learned:
- Array declaration and initialization
- `len()` with arrays
- Accessing elements by index

Challenge:
Find the average of the elements in the array.

---

## Exercise 22 — Slices Basics
Difficulty: Intermediate
Goal: Use the most common sequence type in Go

Description:
Create a slice of strings containing 3 fruit names. Use the `append` function to add 2 more fruits to the slice. Print the resulting slice.

Concepts learned:
- Slice vs. Array
- `append()` function
- Slice literals

Challenge:
Create a new slice by taking a "slice" of the existing one (e.g., `fruits[1:3]`).

---

## Exercise 23 — Slice Length and Capacity
Difficulty: Intermediate
Goal: Understand how slices grow

Description:
Create a slice using `make` with a length of 0 and a capacity of 5. In a loop, append 10 integers to it. In each iteration, print the element, the current length, and the current capacity.

Concepts learned:
- `make()` function
- Length (`len`) vs. Capacity (`cap`)
- How Go reallocates memory for slices

Challenge:
Create a slice of a specific size and see what happens when you try to access an index outside of it.

---

## Exercise 24 — Iterating over Slices (Range)
Difficulty: Intermediate
Goal: Use range for cleaner loops

Description:
Create a slice of integers. Use `for range` to iterate over the slice and print both the index and the value of each element.

Concepts learned:
- `for range` loop
- Ignoring index or value with `_`

Challenge:
Calculate the sum of only the values in the slice using range.

---

## Exercise 25 — Maps Basics
Difficulty: Intermediate
Goal: Use key-value pairs

Description:
Create a map where keys are string names and values are integer ages. Add 3 people to the map and print the age of one specific person.

Concepts learned:
- Map declaration and initialization
- Adding and retrieving map values
- Unordered nature of maps

Challenge:
Check if a key exists in the map before accessing it using the "comma ok" idiom.

---

## Exercise 26 — Deleting from Maps
Difficulty: Intermediate
Goal: Manage map content

Description:
Start with the map from Exercise 25. Delete one person from the map using the `delete` function and print the whole map to verify.

Concepts learned:
- `delete()` function
- Iterating over maps with `range`

Challenge:
Clear all elements from a map (Go 1.21+ has `clear(m)`).

---

## Exercise 27 — Structs
Difficulty: Intermediate
Goal: Group related data

Description:
Define a `Person` struct with `Name` (string), `Age` (int), and `Email` (string). Create an instance of this struct and print its fields.

Concepts learned:
- Struct definition
- Struct instantiation (literal and `new`)
- Dot notation for accessing fields

Challenge:
Create an anonymous struct for a one-time use scenario.

---

## Exercise 28 — Embedded Structs (Composition)
Difficulty: Intermediate
Goal: Use composition instead of inheritance

Description:
Define a `ContactInfo` struct. Then, embed `ContactInfo` inside a `User` struct. Create a `User` and access the embedded fields directly.

Concepts learned:
- Struct embedding
- Field promotion

Challenge:
Define a field with the same name in both the parent and embedded struct. See how to access each.

---

## Exercise 29 — Methods
Difficulty: Intermediate
Goal: Add behavior to types

Description:
Add a method `IsAdult()` to the `Person` struct from Exercise 27 that returns `true` if the person is 18 or older.

Concepts learned:
- Method receivers
- Value vs. Pointer receivers (basics)

Challenge:
Add a `Greet()` method that returns a formatted string including the person's name.

---

## Exercise 30 — Pointer Receivers
Difficulty: Intermediate
Goal: Modify struct state in methods

Description:
Add a method `HaveBirthday()` to the `Person` struct that increments their age by 1. Use a pointer receiver so the change persists.

Concepts learned:
- Pointer receivers (`*`)
- When to use pointer vs. value receivers

Challenge:
Try making it a value receiver and see why the age doesn't change in `main`.

---

## Exercise 31 — Pointers Basics
Difficulty: Intermediate
Goal: Understand memory addresses

Description:
Declare an integer and a pointer to that integer. Print the value of the integer, its memory address, and the value the pointer points to (dereferencing).

Concepts learned:
- Address-of operator (`&`)
- Pointer type (`*int`)
- Dereference operator (`*`)

Challenge:
Write a function `swap(a, b *int)` that swaps the values of two integers.

---

## Exercise 32 — Error Handling (The Go Way)
Difficulty: Intermediate
Goal: Handle errors explicitly

Description:
Write a function `Sqrt(f float64) (float64, error)` that returns the square root of a number. If the number is negative, return an error using `errors.New`.

Concepts learned:
- `error` type
- `errors.New`
- The `if err != nil` pattern

Challenge:
Use `fmt.Errorf` to include the invalid value in the error message.

---

## Exercise 33 — Custom Error Types
Difficulty: Intermediate
Goal: Define rich error information

Description:
Create a custom error type `DivisionError` (a struct) that includes the numerator and denominator. Implement the `Error()` method.

Concepts learned:
- Implementing the `error` interface
- Type assertions for errors

Challenge:
Use `errors.As` (Go 1.13+) to catch your custom error type.

---

## Exercise 34 — Panic and Recover
Difficulty: Intermediate
Goal: Handle catastrophic failures

Description:
Write a program that intentionally triggers a panic (e.g., division by zero or index out of bounds). Use `defer` and `recover` to catch the panic and prevent the program from crashing.

Concepts learned:
- `panic`
- `recover`
- `defer` execution timing

Challenge:
Wrap logic in a function that converts panics into standard errors.

---

## Exercise 35 — Defer Keyword
Difficulty: Intermediate
Goal: Manage resources and execution order

Description:
Write a program that prints "Starting", "Middle", and "Ending". Use `defer` to ensure "Ending" is printed last, even if "Middle" is called before it.

Concepts learned:
- `defer` statement
- LIFO (Last-In-First-Out) order of deferred calls

Challenge:
Use `defer` to time how long a function takes to execute (printing the duration at the end).

---

## Exercise 36 — Packages and Imports
Difficulty: Intermediate
Goal: Organize code into modules

Description:
Create a new directory `calculator`. Inside it, create a file `math.go` in `package calculator` with an `Add` function (capitalized for export). Import this package into your `main.go` and use it.

Concepts learned:
- Package declaration
- Exported vs. Unexported identifiers (Capitalization)
- Local imports

Challenge:
Rename the imported package using an alias (e.g., `import calc "path/to/calculator"`).

---

## Exercise 37 — Go Modules (init)
Difficulty: Intermediate
Goal: Initialize a Go project

Description:
Create a new folder and run `go mod init <module-name>`. Research what the `go.mod` file does.

Concepts learned:
- `go mod init`
- Module path and Go versions in `go.mod`

Challenge:
Add a third-party dependency (like `github.com/google/uuid`) using `go get` and see how `go.mod` and `go.sum` change.

---

## Exercise 38 — init() Function
Difficulty: Intermediate
Goal: Perform package-level initialization

Description:
Write a program with an `init()` function that sets up some configuration (e.g., setting a global variable) before `main()` runs.

Concepts learned:
- Execution order of `init()` vs `main()`
- Multiple `init()` functions (optional)

Challenge:
Check how `init()` functions run across multiple imported packages.

---

## Exercise 39 — JSON Serialization (Struct Tags)
Difficulty: Intermediate
Goal: Convert structs to JSON

Description:
Define a `Product` struct with fields for Name, Price, and Stock. Use struct tags to make the JSON keys lowercase. Marshal a `Product` instance into a JSON string and print it.

Concepts learned:
- `encoding/json` package
- Struct tags (`` `json:"name"` ``)
- `json.Marshal`

Challenge:
Write a program that takes a slice of Products and marshals it into a JSON array.

---

# Phase 3 — Go Core Concepts (Exercises 41–60)

## Exercise 41 — Interfaces Basics
Difficulty: Intermediate
Goal: Define behavior through interfaces

Description:
Define a `Shape` interface with an `Area() float64` method. Create `Circle` and `Rectangle` structs that implement this interface.

Concepts learned:
- Interface definition
- Implicit interface implementation
- Polymorphism in Go

Challenge:
Write a function `PrintArea(s Shape)` that prints the area of any shape passed to it.

---

## Exercise 42 — Empty Interface (any)
Difficulty: Intermediate
Goal: Work with values of unknown types

Description:
Create a function that takes an `interface{}` (or `any` in Go 1.18+) as a parameter and prints its value and type using `%v` and `%T`.

Concepts learned:
- `interface{}` / `any`
- Type flexibility

Challenge:
Use a type switch to perform different actions based on whether the input is an `int`, `string`, or `bool`.

---

## Exercise 43 — Type Assertions
Difficulty: Intermediate
Goal: Extract underlying values from interfaces

Description:
Create an interface variable holding a string. Use a type assertion to extract the string value. Handle the case where the assertion might fail using the "comma ok" idiom.

Concepts learned:
- Type assertion syntax `v := i.(type)`
- Comma-ok idiom for safe assertions

Challenge:
Try to assert an interface holding an `int` to a `float64` and observe the behavior.

---

## Exercise 44 — Stringer Interface
Difficulty: Intermediate
Goal: Customize how types are printed

Description:
Implement the `fmt.Stringer` interface for the `Person` struct from Phase 2. The `String()` method should return "Name (Age)".

Concepts learned:
- `fmt.Stringer` interface
- Customizing string output for types

Challenge:
Implement a custom stringer for a slice of `Person` that returns a comma-separated list.

---

## Exercise 45 — Readers and Writers
Difficulty: Intermediate
Goal: Work with I/O streams

Description:
Use `strings.NewReader` to create a reader from a string. Use `io.Copy` to copy the content of the reader to `os.Stdout`.

Concepts learned:
- `io.Reader` and `io.Writer` interfaces
- `io.Copy`
- `strings.Reader`

Challenge:
Write a custom `io.Writer` that counts how many bytes are written to it.

---

## Exercise 46 — Reading Files (os.ReadFile)
Difficulty: Intermediate
Goal: Read simple files

Description:
Write a program that reads the entire content of a text file using `os.ReadFile` and prints it to the console.

Concepts learned:
- `os.ReadFile`
- Error handling for file operations

Challenge:
Check if the file exists before attempting to read it using `os.Stat`.

---

## Exercise 47 — Writing Files (os.WriteFile)
Difficulty: Intermediate
Goal: Write simple files

Description:
Write a program that takes user input and saves it to a file named `output.txt` using `os.WriteFile`.

Concepts learned:
- `os.WriteFile`
- File permissions (e.g., `0644`)

Challenge:
Append text to an existing file instead of overwriting it using `os.OpenFile`.

---

## Exercise 48 — Buffered File I/O (bufio)
Difficulty: Intermediate
Goal: Efficiently read/write large files

Description:
Use `bufio.Scanner` to read a file line by line. Count the number of lines in the file.

Concepts learned:
- `bufio.Scanner`
- Iterating over file lines efficiently

Challenge:
Search for a specific word in a file and print the line numbers where it occurs.

---

## Exercise 49 — Creating CLI Tools (Flags)
Difficulty: Intermediate
Goal: Parse command-line arguments

Description:
Create a CLI tool that takes a `-name` string flag and a `-age` int flag. Print a greeting using these values.

Concepts learned:
- `flag` package
- `flag.Parse()`
- Pointer flags vs value flags

Challenge:
Add a boolean flag `-verbose` that, when set, prints extra debug information.

---

## Exercise 50 — Subcommands in CLI
Difficulty: Intermediate
Goal: Build complex CLI structures

Description:
Use the `flag` package to create a tool with two subcommands: `greet` (takes a name) and `time` (prints current time).

Concepts learned:
- `flag.NewFlagSet`
- Subcommand routing logic

Challenge:
Add a `help` command that explains how to use the tool.

---

## Exercise 51 — Unit Testing Basics
Difficulty: Intermediate
Goal: Write automated tests

Description:
Create a file `math.go` with a `Sum(a, b int) int` function. Create a corresponding `math_test.go` file and write a test function `TestSum` to verify its correctness.

Concepts learned:
- `testing` package
- Test file naming convention (`_test.go`)
- `t.Errorf` for failures

Challenge:
Run your tests using `go test -v`.

---

## Exercise 52 — Table-Driven Tests
Difficulty: Intermediate
Goal: Write concise and thorough tests

Description:
Rewrite the test in Exercise 51 as a table-driven test using a slice of anonymous structs containing inputs and the expected output.

Concepts learned:
- Table-driven testing pattern
- Iterating over test cases
- Subtests with `t.Run`

Challenge:
Add edge cases to your test table (e.g., negative numbers, zero).

---

## Exercise 53 — Benchmarking
Difficulty: Intermediate
Goal: Measure code performance

Description:
Write a benchmark function `BenchmarkSum` for your `Sum` function.

Concepts learned:
- `testing.B`
- Running benchmarks with `go test -bench=.`

Challenge:
Compare the performance of two different ways of reversing a string.

---

## Exercise 54 — Sorting Basics
Difficulty: Intermediate
Goal: Use the sort package

Description:
Create a slice of integers and sort them in ascending order using the `sort` package.

Concepts learned:
- `sort.Ints`
- `sort.Strings`

Challenge:
Sort a slice of `Person` structs by age using `sort.Slice`.

---

## Exercise 55 — Working with Time
Difficulty: Intermediate
Goal: Handle dates and times

Description:
Write a program that prints the current date and time in a custom format (e.g., "YYYY-MM-DD HH:MM:SS").

Concepts learned:
- `time` package
- `time.Now()`
- Time formatting with the "2006-01-02 15:04:05" layout

Challenge:
Calculate the duration between two specific dates (e.g., how many days until your next birthday).

---

## Exercise 56 — Environment Variables
Difficulty: Intermediate
Goal: Configure apps using the environment

Description:
Write a program that reads an environment variable `APP_PORT`. If it's not set, use a default value of `8080`.

Concepts learned:
- `os.Getenv`
- Configuration defaults

Challenge:
Use `os.Setenv` within a test to verify your program's behavior with different environment values.

---

## Exercise 57 — Working with URLs
Difficulty: Intermediate
Goal: Parse and build URLs

Description:
Take a URL string (e.g., "https://example.com/search?q=golang") and parse it using the `net/url` package. Print the scheme, host, path, and query parameters.

Concepts learned:
- `net/url` package
- `url.Parse`
- Accessing query params

Challenge:
Construct a new URL string programmatically from parts.

---

## Exercise 58 — Basic Reflection (Type and Value)
Difficulty: Advanced
Goal: Inspect types at runtime

Description:
Use the `reflect` package to print the `Type` and `Value` of an unknown variable.

Concepts learned:
- `reflect.TypeOf`
- `reflect.ValueOf`
- Reflection dangers/performance

Challenge:
Write a function that prints all field names and values of an arbitrary struct passed to it using reflection.

---

## Exercise 59 — Embedding Interfaces
Difficulty: Intermediate
Goal: Compose behaviors

Description:
Define interfaces `Reader` and `Writer`. Then define a `ReadWriter` interface that embeds both. Implement a struct that satisfies `ReadWriter`.

Concepts learned:
- Interface embedding
- Composite interfaces

Challenge:
Use the standard library `io.ReadWriter` in a function.

---

# Phase 4 — Concurrency & Performance (Exercises 61–80)

## Exercise 61 — My First Goroutine
Difficulty: Intermediate
Goal: Understand basic concurrency

Description:
Write a program that launches a goroutine to print "Hello from goroutine" while the main function prints "Hello from main". Use `time.Sleep` to ensure the goroutine has time to finish before the program exits.

Concepts learned:
- `go` keyword
- Main goroutine behavior
- Why `time.Sleep` is usually a bad way to wait for goroutines

Challenge:
Launch 10 goroutines and see the order of their execution.

---

## Exercise 62 — WaitGroups
Difficulty: Intermediate
Goal: Synchronize goroutines properly

Description:
Use `sync.WaitGroup` to wait for 3 goroutines to finish their execution. Each goroutine should print a message and then call `Done()`.

Concepts learned:
- `sync.WaitGroup`
- `Add`, `Done`, and `Wait` methods
- Passing pointers to WaitGroups

Challenge:
What happens if you call `Add(3)` but only call `Done()` twice?

---

## Exercise 63 — Unbuffered Channels
Difficulty: Intermediate
Goal: Communicate between goroutines

Description:
Create an unbuffered channel of integers. Launch a goroutine that sends a value into the channel, and receive that value in the `main` goroutine.

Concepts learned:
- Channel declaration (`chan int`)
- Sending (`ch <- v`) and receiving (`v := <-ch`)
- Channel blocking behavior

Challenge:
Try to receive from a channel when no one is sending. What error do you get?

---

## Exercise 64 — Buffered Channels
Difficulty: Intermediate
Goal: Understand channel capacity

Description:
Create a buffered channel with a capacity of 2. Send two values into it without a receiver, then receive them both.

Concepts learned:
- Buffered channels (`make(chan int, 2)`)
- Differences in blocking between buffered and unbuffered channels

Challenge:
Send 3 values into a channel with capacity 2 without receiving. Observe the deadlock.

---

## Exercise 65 — Channel Closing and Range
Difficulty: Intermediate
Goal: Signal completion via channels

Description:
Create a goroutine that sends the first 5 numbers of the Fibonacci sequence into a channel and then closes it. Use a `for range` loop in `main` to receive all values.

Concepts learned:
- `close(ch)`
- `for range` over channels
- Receiving from a closed channel

Challenge:
Use the "comma ok" idiom to check if a channel is closed while receiving.

---

## Exercise 66 — Select Statement
Difficulty: Intermediate
Goal: Handle multiple channel operations

Description:
Create two channels and two goroutines that send values at different intervals (use `time.Sleep`). Use a `select` statement in `main` to print values as they arrive.

Concepts learned:
- `select` statement
- Multiplexing channels

Challenge:
Add a `default` case to the `select` and a timeout using `time.After`.

---

## Exercise 67 — Mutual Exclusion (Mutex)
Difficulty: Intermediate
Goal: Prevent race conditions

Description:
Create a global counter variable. Launch 1000 goroutines that each increment the counter. Use `sync.Mutex` to ensure the final count is exactly 1000.

Concepts learned:
- Race conditions
- `sync.Mutex`
- `Lock` and `Unlock`

Challenge:
Run the program with the race detector: `go run -race main.go`.

---

## Exercise 68 — Read-Write Mutex (RWMutex)
Difficulty: Intermediate
Goal: Optimize for frequent reads

Description:
Create a shared map protected by a `sync.RWMutex`. Launch many "reader" goroutines and a few "writer" goroutines. Note the performance difference compared to a standard `Mutex`.

Concepts learned:
- `sync.RWMutex`
- `RLock` vs `Lock`

Challenge:
Simulate a "cache" where many goroutines read a value and only one updates it occasionally.

---

## Exercise 69 — Atomic Operations
Difficulty: Intermediate
Goal: High-performance synchronization for simple types

Description:
Use the `sync/atomic` package to increment a counter from multiple goroutines instead of using a Mutex.

Concepts learned:
- `sync/atomic` package
- `atomic.AddInt64`, `atomic.LoadInt64`

Challenge:
Compare the benchmark of Mutex vs Atomic for a simple counter increment.

---

## Exercise 70 — Worker Pool Pattern
Difficulty: Intermediate
Goal: Limit concurrency and distribute tasks

Description:
Implement a worker pool with 3 workers. Create a `jobs` channel and a `results` channel. Each worker should read from `jobs`, process the task (e.g., multiply by 2), and send to `results`.

Concepts learned:
- Worker pool architecture
- Distributing work over channels
- Collecting results

Challenge:
Dynamically increase or decrease the number of workers based on demand.

---

## Exercise 71 — Generator Pattern
Difficulty: Intermediate
Goal: Produce streams of data

Description:
Write a function `fibonacci()` that returns a receive-only channel (`<-chan int`) which generates Fibonacci numbers indefinitely.

Concepts learned:
- Directional channels
- Generator pattern

Challenge:
Add a `quit` channel to the generator to stop it gracefully.

---

## Exercise 72 — Fan-in Pattern
Difficulty: Intermediate
Goal: Combine multiple channels into one

Description:
Write a function `FanIn(ch1, ch2 <-chan string) <-chan string` that merges two channels into a single output channel.

Concepts learned:
- Fan-in pattern
- Using goroutines within functions to bridge channels

Challenge:
Implement Fan-in for an arbitrary number of input channels (variadic function).

---

## Exercise 73 — Fan-out Pattern
Difficulty: Intermediate
Goal: Distribute one channel to many workers

Description:
Take a single channel of tasks and distribute them among multiple worker goroutines.

Concepts learned:
- Fan-out pattern
- Load balancing basics

Challenge:
Combine Fan-out and Fan-in to process data in parallel and collect results.

---

## Exercise 74 — Context Package (Cancellation)
Difficulty: Intermediate
Goal: Manage goroutine lifecycles

Description:
Create a goroutine that performs a long-running task. Use `context.WithCancel` to stop the goroutine from `main` after 2 seconds.

Concepts learned:
- `context` package
- `ctx.Done()`
- Propagating cancellation

Challenge:
Handle multiple nested goroutines and ensure they all stop when the parent context is cancelled.

---

## Exercise 75 — Context with Timeout
Difficulty: Intermediate
Goal: Prevent hanging operations

Description:
Use `context.WithTimeout` to limit the execution time of a function. If the timeout is reached, the function should return an error.

Concepts learned:
- `context.WithTimeout`
- Handling `context.DeadlineExceeded`

Challenge:
Use context timeout with an HTTP request (using `http.NewRequestWithContext`).

---

## Exercise 76 — Pipeline Pattern
Difficulty: Intermediate
Goal: Chain processing stages

Description:
Build a multi-stage pipeline:
1. Stage 1: Generate integers.
2. Stage 2: Square the integers.
3. Stage 3: Print the integers.
Each stage should be a function connected by channels.

Concepts learned:
- Pipeline pattern
- Separation of concerns in processing

Challenge:
Add a stage that filters out odd numbers.

---

## Exercise 77 — Rate Limiting
Difficulty: Intermediate
Goal: Control the rate of operations

Description:
Use a `time.Ticker` to implement a simple rate limiter that allows only 5 requests per second.

Concepts learned:
- `time.Ticker`
- Tick-based synchronization

Challenge:
Implement a "bursty" rate limiter that allows a few requests immediately but limits the sustained rate.

---

## Exercise 78 — Once Initialization
Difficulty: Intermediate
Goal: Perform expensive setup exactly once

Description:
Use `sync.Once` to ensure that a configuration setup function is only ever executed once, even if called from multiple goroutines.

Concepts learned:
- `sync.Once`
- Thread-safe singleton patterns

Challenge:
Try to implement the same logic using just a boolean and a Mutex, then compare the complexity.

---

## Exercise 79 — Performance Profiling (pprof)
Difficulty: Advanced
Goal: Identify bottlenecks

Description:
Create a program with a known performance issue (e.g., excessive allocations). Use the `runtime/pprof` package to generate a CPU profile and analyze it using `go tool pprof`.

Concepts learned:
- CPU and Memory profiling
- `go tool pprof`
- Identifying "hot spots" in code

Challenge:
Optimize the code based on the profile results and verify the improvement with another profile.

---

# Phase 5 — Real Backend Development (Exercises 81–100)

## Exercise 81 — Simple HTTP Server
Difficulty: Intermediate
Goal: Serve web requests

Description:
Use the `net/http` package to create a simple web server that listens on port 8080 and returns "Welcome to Go Backend!" when the root path (`/`) is visited.

Concepts learned:
- `http.HandleFunc`
- `http.ListenAndServe`
- Writing to `http.ResponseWriter`

Challenge:
Add a dynamic route like `/hello?name=John` that greets the user by name.

---

## Exercise 82 — JSON API Endpoint
Difficulty: Intermediate
Goal: Return structured data

Description:
Create an HTTP handler that returns a JSON object representing a User (from Exercise 27). Set the `Content-Type` header to `application/json`.

Concepts learned:
- Encoding JSON directly to the response writer
- Setting HTTP headers

Challenge:
Return a slice of Users as a JSON array.

---

## Exercise 83 — HTTP Methods (POST)
Difficulty: Intermediate
Goal: Handle data submission

Description:
Create a POST endpoint `/users` that accepts a JSON body to create a new user. Log the user details to the console.

Concepts learned:
- Checking `r.Method`
- Decoding JSON from `r.Body`

Challenge:
Validate that the user's email is not empty before "creating" them.

---

## Exercise 84 — Serving Static Files
Difficulty: Intermediate
Goal: Serve HTML/CSS/JS

Description:
Use `http.FileServer` to serve static files from a directory named `static`. Place an `index.html` file there and verify it can be accessed.

Concepts learned:
- `http.FileServer`
- `http.Dir`

Challenge:
Combine static file serving with dynamic API routes.

---

## Exercise 85 — Custom Mux (http.ServeMux)
Difficulty: Intermediate
Goal: Organize routes properly

Description:
Instead of using the default mux, create your own `http.ServeMux` and register handlers on it.

Concepts learned:
- `http.NewServeMux()`
- Modularizing routing

Challenge:
Create a "Router" struct that wraps the mux and provides helper methods for registration.

---

## Exercise 86 — Middleware Basics
Difficulty: Intermediate
Goal: Intercept and process requests

Description:
Write a simple middleware function that logs the HTTP method and path of every incoming request before calling the next handler.

Concepts learned:
- Middleware pattern in Go
- Wrapping `http.Handler`

Challenge:
Create a middleware that adds a custom header (e.g., `X-Request-ID`) to every response.

---

## Exercise 87 — Recovery Middleware
Difficulty: Intermediate
Goal: Prevent server crashes

Description:
Implement a middleware that uses `recover()` to catch any panics that happen inside your handlers and returns a 500 Internal Server Error instead of crashing the process.

Concepts learned:
- Panic recovery in HTTP servers
- Graceful error responses

Challenge:
Log the stack trace of the panic to the console.

---

## Exercise 88 — Working with Chi or Gorilla Mux
Difficulty: Intermediate
Goal: Use a popular third-party router

Description:
Initialize a new Go module and install a popular router library like `github.com/go-chi/chi`. Rewrite your previous API using Chi's routing features.

Concepts learned:
- Third-party routing libraries
- URL parameters (e.g., `/users/{id}`)

Challenge:
Implement route grouping for a specific path prefix (e.g., `/api/v1`).

---

## Exercise 89 — Environment Configuration (Godotenv)
Difficulty: Intermediate
Goal: Load config from .env files

Description:
Use the `github.com/joho/godotenv` library to load environment variables from a `.env` file into your application.

Concepts learned:
- Configuration management
- .env files vs system envs

Challenge:
Create a `Config` struct that holds all your app's settings and populate it on startup.

---

## Exercise 90 — Connecting to PostgreSQL
Difficulty: Advanced
Goal: Database integration

Description:
Set up a local PostgreSQL database. Use the `database/sql` package and a driver like `github.com/lib/pq` or `github.com/jackc/pgx` to connect to it.

Concepts learned:
- `sql.Open`
- Connection strings
- Pinging the database

Challenge:
Implement a connection pool with set limits on max open connections.

---

## Exercise 91 — Running SQL Queries
Difficulty: Advanced
Goal: Execute basic DB operations

Description:
Create a `users` table in PostgreSQL. Write a Go program that inserts a new user into the table and then retrieves it by ID.

Concepts learned:
- `db.Exec` for writes
- `db.Query` and `db.QueryRow` for reads
- Scanning results into structs

Challenge:
Use `sql.NullString` to handle optional columns that might be NULL in the database.

---

## Exercise 92 — Using SQLX
Difficulty: Advanced
Goal: Simplify database interaction

Description:
Install `github.com/jmoiron/sqlx`. Rewrite your database queries to use `sqlx`'s struct tagging and `Get`/`Select` methods.

Concepts learned:
- `sqlx` vs raw `database/sql`
- Mapping DB columns to struct fields automatically

Challenge:
Implement a query that uses `IN` with a slice of IDs.

---

## Exercise 93 — Database Migrations
Difficulty: Advanced
Goal: Manage schema changes

Description:
Use a tool like `golang-migrate` or `go-migrate` (or write your own simple script) to manage your database schema via migration files.

Concepts learned:
- Schema versioning
- Up/Down migrations

Challenge:
Automate migrations to run whenever the server starts.

---

## Exercise 94 — JWT Authentication
Difficulty: Advanced
Goal: Secure your API

Description:
Implement a login endpoint that verifies a password (use a hardcoded one for now) and returns a JSON Web Token (JWT) using `github.com/golang-jwt/jwt`.

Concepts learned:
- JWT structure (Header, Payload, Signature)
- Signing tokens

Challenge:
Implement an "Auth" middleware that validates the JWT in the `Authorization` header for protected routes.

---

## Exercise 95 — Password Hashing (Bcrypt)
Difficulty: Intermediate
Goal: Store passwords securely

Description:
Use `golang.org/x/crypto/bcrypt` to hash a user's password before storing it and to verify a password against a hash during login.

Concepts learned:
- Salting and hashing
- Bcrypt costs

Challenge:
Implement a full "Register" flow where the user's password is hashed before being saved to the DB.

---

## Exercise 96 — Structured Logging (Slog)
Difficulty: Intermediate
Goal: Better production logs

Description:
Use the standard library `log/slog` (available in Go 1.21+) to implement structured logging in your application.

Concepts learned:
- Key-value logging
- Log levels (Info, Debug, Error)
- JSON log formatters

Challenge:
Update your middleware to include the request duration in the structured logs.

---

## Exercise 97 — Dependency Injection (Manual)
Difficulty: Intermediate
Goal: Improve testability and decoupling

Description:
Refactor your server so that the Handlers receive the Database connection as a dependency (constructor injection) rather than using a global variable.

Concepts learned:
- Dependency injection patterns
- Struct-based handlers

Challenge:
Create a "Service" layer that sits between your handlers and the database.

---

## Exercise 98 — Integration Testing
Difficulty: Advanced
Goal: Test the whole system

Description:
Write an integration test that starts an instance of your server and makes actual HTTP requests to it using `net/http/httptest`.

Concepts learned:
- `httptest.NewServer`
- Testing API endpoints from the outside

Challenge:
Use a "test database" that is cleaned up before or after each test run.

---

## Exercise 99 — Graceful Shutdown
Difficulty: Advanced
Goal: Stop servers safely

Description:
Modify your server to listen for OS signals (like SIGINT/SIGTERM). When a signal is received, wait for active requests to finish before shutting down the server.

Concepts learned:
- Signal handling with `os/signal`
- `server.Shutdown()` with context

Challenge:
Close database connections and other resources during the shutdown phase.

---

## Exercise 100 — Production-Ready Service
Difficulty: Advanced
Goal: Final Capstone Project

Description:
Build a "Task Manager" API that combines everything you've learned:
- CRUD operations for Tasks (ID, Title, Done) stored in PostgreSQL.
- JWT-based Authentication.
- Structured logging.
- Error handling middleware.
- Environment configuration.
- Dependency injection.
- Integration tests.
- Graceful shutdown.

Concepts learned:
- Architecture of a real-world Go service
- Project organization and best practices

Challenge:
Containerize your application and database using Docker and Docker Compose.
