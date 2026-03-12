# Day 46 - Reading Files (os.ReadFile)

## Goal
Read simple files.

## Simple Explanation
For small files, Go provides a simple way to read the whole thing at once using `os.ReadFile`. This function takes the file path and returns the entire content as a slice of bytes (`[]byte`). You can then easily convert those bytes into a string to read it like normal text.

## Problem Description
Write a program that reads the entire content of a text file using `os.ReadFile` and prints it to the console.
-
