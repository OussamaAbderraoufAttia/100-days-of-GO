# Day 62 - WaitGroups (sync package)

## Goal
Properly wait for goroutines to finish.

## Simple Explanation
Using `time.Sleep` to wait for goroutines is not reliable because we don't know exactly how long they will take. instead, we use a **WaitGroup**. It's like a counter:
1. `Add(1)`: Tell the counter we started a new job.
2. `Done()`: The goroutine calls this when it finishes to subtract 1.
3. `Wait()`: The main program pauses here until the counter reaches zero.

## Problem Description
Use `sync.WaitGroup` to wait for three separate goroutines to complete their work.
-
