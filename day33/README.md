# Day 33 - Custom Error Types

## Goal
Define rich error information.

## Simple Explanation
Sometimes, a simple text message isn't enough to explain an error. We can create our own `error` type by defining a struct and adding an `Error() string` method to it. This lets us include more details, like the data that caused the error, which can be very helpful for debugging.

## Problem Description
Create a custom error type `DivisionError` (a struct) that includes the numerator and denominator. Implement the `Error()` method.
