# Day 61 - Goroutines (The `go` keyword)

## Goal
Run code concurrently.

## Simple Explanation
A **Goroutine** is a lightweight thread managed by the Go runtime. When you put the word `go` in front of a function call, Go starts that function in the background and carries on with the rest of your program immediately. This lets you do many things at the same time without waiting for one to finish before starting the next.

## Problem Description
Create a function that prints numbers 1 to 5 with a small delay. Call it as a goroutine and observe that the program might exit before the goroutine finishes unless you add a temporary delay in `main`.
-
