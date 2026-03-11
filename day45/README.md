# Day 45 - Readers and Writers

## Goal
Work with I/O streams.

## Simple Explanation
In Go, `io.Reader` and `io.Writer` are two of the most important interfaces. A **Reader** is anything you can read bytes from (like a file, a network connection, or a string). A **Writer** is anything you can write bytes into (like the console, a file, or an HTTP response). This standard way of handling data makes it incredibly easy to "pipe" data from one place to another.

## Problem Description
Use `strings.NewReader` to create a reader from a string. Use `io.Copy` to copy the content of the reader to `os.Stdout`.
