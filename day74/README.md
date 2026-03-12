# Day 74 - Sync.Once

## Goal
Ensure a function runs only once.

## Simple Explanation
`sync.Once` is a special Go tool that guarantees a piece of code runs exactly one time, no matter how many goroutines try to call it. It's perfect for things like loading configurations or connecting to a database for the first time. Even if 100 workers try to "init" at the same moment, Go ensures it only happens once and the others wait until it's finished.

## Problem Description
Create 10 goroutines that all attempt to call an `initialize()` function. Use `sync.Once` to ensure the message "Initializing..." is only printed once.
-
