# Day 48 - Buffered File I/O (bufio)

## Goal
Efficiently read/write large files.

## Simple Explanation
When files are very large, reading them all at once (like `ReadFile`) might use up all your computer's memory. Instead, we use a **Scanner** from the `bufio` package to read the file line-by-line. This is very efficient because Go only keeps one small part of the file in memory at a time.

## Problem Description
Use `bufio.Scanner` to read a file line by line. Count the number of lines in the file.
-
