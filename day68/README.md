# Day 68 - Atomic Operations (sync/atomic)

## Goal
Perform simple thread-safe math.

## Simple Explanation
For simple cases like counting numbers, using a Mutex can be a bit slow. Go provides the `sync/atomic` package, which uses special computer hardware instructions to update numbers in a way that is guaranteed to be safe from multiple workers. It's faster than a Mutex but only works for basic math on integers and pointers.

## Problem Description
Use the `sync/atomic` package to create a concurrent counter without using Mutexes.
