# Day 79 - Errors with Concurrency (errgroup)

## Goal
Handle errors from many goroutines.

## Simple Explanation
WaitGroups are great for waiting, but they don't help if one of the workers hits an error. The `errgroup` package (part of Go's extra tools) is like a WaitGroup that also gathers errors. If even one worker fails, `errgroup` will catch it and tell you about it, making it much safer for complex tasks.

## Problem Description
Research the `golang.org/x/sync/errgroup` package. Write a conceptual program that uses it to run three tasks and catches the first error that occurs.
-
