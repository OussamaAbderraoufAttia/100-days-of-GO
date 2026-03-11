# Day 34 - Panic and Recover

## Goal
Handle catastrophic failures.

## Simple Explanation
`panic` is a way to stop the program immediately when an unrecoverable error happens (like a bug). `recover` is a special function that can stop the panic and let the program continue. We usually use `recover` inside a `defer` function to catch unexpected crashes, though we should only use it for truly exceptional cases.

## Problem Description
Write a program that intentionally triggers a panic (e.g., division by zero or index out of bounds). Use `defer` and `recover` to catch the panic and prevent the program from crashing.
