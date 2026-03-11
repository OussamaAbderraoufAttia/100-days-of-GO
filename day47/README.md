# Day 47 - Writing Files (os.WriteFile)

## Goal
Write simple files.

## Simple Explanation
Just like reading, writing small files is very easy with `os.WriteFile`. You give it the file name, the data (as bytes), and "file permissions" (numbers that tell the computer who can read or change the file). If the file doesn't exist, Go will create it for you.

## Problem Description
Write a program that takes user input and saves it to a file named `output.txt` using `os.WriteFile`.
