# Day 49 - Creating CLI Tools (Flags)

## Goal
Parse command-line arguments.

## Simple Explanation
Many Go programs are "command-line tools" meant to be run from the terminal. The `flag` package lets us easily accept inputs when the program starts, like `myprogram -name Al -age 25`. Each "flag" is like a variable that the user can set by typing its name in the terminal.

## Problem Description
Create a CLI tool that takes a `-name` string flag and a `-age` int flag. Print a greeting using these values.
