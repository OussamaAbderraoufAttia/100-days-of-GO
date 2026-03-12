# Day 53 - Benchmarking

## Goal
Measure code performance.

## Simple Explanation
Benchmarking tells you how *fast* your code is. In Go, you write functions starting with `Benchmark` and use `testing.B`. Go will run your code thousands or millions of times to find out exactly how many nanoseconds it takes on average. This is great for making sure your changes don't slow down the program.

## Problem Description
Write a benchmark function `BenchmarkSum` for your `Sum` function.
-
