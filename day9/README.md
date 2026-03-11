# Day 9 - Infinite Loops and Break

## Goal
Control flow with break and continue.

## Simple Explanation
Creating an infinite loop is easy in Go: just type `for { ... }`. To stop it, we use the `break` command. We can also use `continue` to jump immediately to the next turn of the loop, skipping the code after the `continue` command.

## Problem Description
Write an infinite loop that increments a counter. Inside the loop, use an `if` statement to `break` the loop when the counter reaches 5. Use `continue` to skip printing the number 3.
