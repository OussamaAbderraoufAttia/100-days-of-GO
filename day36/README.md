# Day 36 - Packages and Imports

## Goal
Organize code into modules.

## Simple Explanation
As your code grows, you shouldn't put everything in one file. **Packages** let you group related code together. You can "export" a function (make it usable by other packages) by starting its name with a Capital letter. Small names for packages are best.

## Problem Description
Create a new directory `calculator`. Inside it, create a file `math.go` in `package calculator` with an `Add` function (capitalized for export). Import this package into your `main.go` and use it.
-
