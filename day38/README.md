# Day 38 - init() Function

## Goal
Perform package-level initialization.

## Simple Explanation
Go has a special function called `init()`. It runs automatically before the `main()` function. You can have many `init()` functions across different files. They are perfect for setting up configurations, connecting to databases, or checking environment variables before your program starts doing its main work.

## Problem Description
Write a program with an `init()` function that sets up some configuration (e.g., setting a global variable) before `main()` runs.
