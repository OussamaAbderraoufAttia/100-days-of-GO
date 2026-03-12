# Day 51 - Unit Testing Basics

## Goal
Write automated tests.

## Simple Explanation
Testing helps us make sure our code works correctly. In Go, test files end with `_test.go`. We use the `testing` package and write functions starting with `Test`. If something is wrong, we use `t.Errorf` to explain what happened. Running `go test` in the terminal runs all these checks for us!

## Problem Description
Create a file `math.go` with a `Sum(a, b int) int` function. Create a corresponding `math_test.go` file and write a test function `TestSum` to verify its correctness.
-
