# Day 13 - User Input (fmt.Scan)

## Goal
Capture input from the console.

## Simple Explanation
To get information from the user, we use `fmt.Scan`. We need to use an ampersand `&` before the variable name. This tells Go to put the user's input directly into that memory address (this is a simple use of "pointers"). Note that `Scan` only reads until it sees a space.

## Problem Description
Write a program that asks the user for their name and age using `fmt.Scan`. Print a greeting that includes these details.
