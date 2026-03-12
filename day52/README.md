# Day 52 - Table-Driven Tests

## Goal
Write concise and thorough tests.

## Simple Explanation
Instead of writing a new test function for every small change, we use **Table-Driven Tests**. We create a "table" (a slice of structs) that lists many different inputs and their expected outputs. We then loop through this table. This makes it very easy to test many "edge cases" (like negative numbers or zero) all at once.

## Problem Description
Rewrite the test in Exercise 51 as a table-driven test using a slice of anonymous structs containing inputs and the expected output.
-
