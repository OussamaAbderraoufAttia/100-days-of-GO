# Day 42 - Empty Interface (any)

## Goal
Work with values of unknown types.

## Simple Explanation
The **Empty Interface** (`interface{}` or simply `any` in newer Go versions) is a special interface that lists zero methods. Because every type has at least zero methods, it means *every* type satisfies the empty interface. This lets us create functions or variables that can hold absolutely any kind of data—strings, integers, structs, you name it!

## Problem Description
Create a function that takes an `interface{}` (or `any` in Go 1.18+) as a parameter and prints its value and type using `%v` and `%T`.
-
