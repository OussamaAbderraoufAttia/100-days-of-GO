# Day 44 - Stringer Interface

## Goal
Customize how types are printed.

## Simple Explanation
Whenever you use `fmt.Println` on an object, Go checks if that object has a `String()` method. This behavior is defined in the `fmt.Stringer` interface. By adding our own `String()` method, we can control exactly how our custom structs look when printed as text, making our logs and output much prettier!

## Problem Description
Implement the `fmt.Stringer` interface for the `Person` struct. The `String()` method should return "Name (Age)".
