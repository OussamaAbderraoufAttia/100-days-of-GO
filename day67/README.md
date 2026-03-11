# Day 67 - Mutexes (sync.Mutex)

## Goal
Protect shared variables.

## Simple Explanation
When multiple goroutines try to change the same variable at the exact same time, it can cause a "Race Condition" (data gets corrupted). A **Mutex** (Mutual Exclusion) is like a lock. A goroutine "Locks" the mutex before changing the data and "Unlocks" it when finished. This ensures only one worker can touch the variable at a time.

## Problem Description
Create a counter variable shared by 1000 goroutines. Use `sync.Mutex` to ensure the final count is exactly 1000.
