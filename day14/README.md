# Day 14 - More Complex Input (bufio)

## Goal
Read full lines of input.

## Simple Explanation
Sometimes we want to read an entire sentence, but `fmt.Scan` stops at the first space. To read a full line, we use the `bufio` package. We create a "Reader" that looks at `os.Stdin` (the keyboard input) and tell it to read until the "Enter" key is pressed (`\n`).

## Problem Description
`fmt.Scan` stops at spaces. Use the `bufio` package and `os.Stdin` to read a full line of text (including spaces) from the user.
-
