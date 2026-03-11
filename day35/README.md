# Day 35 - Defer Keyword

## Goal
Manage resources and execution order.

## Simple Explanation
The `defer` keyword tells Go to wait and run a function right before the *current* function finishes. It's like a scientific cleanup crew—it doesn't matter when you call it, it will always be the last thing to happen. This is perfect for closing files or databases so you never forget to clean up.

## Problem Description
Write a program that prints "Starting", "Middle", and "Ending". Use `defer` to ensure "Ending" is printed last, even if "Middle" is called before it.
