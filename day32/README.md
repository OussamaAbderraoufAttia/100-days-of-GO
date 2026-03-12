# Day 32 - Error Handling (The Go Way)

## Goal
Handle errors explicitly.

## Simple Explanation
In Go, errors are just values that functions return. We don't use "try-catch" blocks. Instead, we check if the returned `error` value is `nil`. If it's not `nil`, something went wrong. This forces us to handle errors immediately and makes the code easier to follow.

## Problem Description
Write a function `Sqrt(f float64) (float64, error)` that returns the square root of a number. If the number is negative, return an error using `errors.New`.
-
