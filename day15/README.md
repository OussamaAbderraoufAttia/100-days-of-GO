# Day 15 - Type Conversion

## Goal
Convert between different data types.

## Simple Explanation
Go is very strict about types. You cannot add an `int` to a `float64` directly. You must manually convert one to the other. To do this, you put the type you want in front of the variable in parentheses, like `int(myFloat)`. Note that converting a decimal to an integer will cut off everything after the decimal point!

## Problem Description
Write a program that takes a `float64` value and converts it to an `int`. Also, convert an `int` to a `float64`. Note how the values change (truncation).
