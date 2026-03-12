# Day 29 - Methods

## Goal
Add behavior to types.

## Simple Explanation
Methods are like functions, but they are "attached" to a specific type (like a struct). When you define a method, you specify a **receiver** in front of the name. For example, `func (p Person) Greet()` means the `Greet` method belongs to `Person`. You can then call it using `myPerson.Greet()`.

## Problem Description
Add a method `IsAdult()` to the `Person` struct from Exercise 27 that returns `true` if the person is 18 or older.
-
