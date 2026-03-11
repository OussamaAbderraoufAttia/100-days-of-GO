# Day 31 - Pointers Basics

## Goal
Understand memory addresses.

## Simple Explanation
A variable is a name for a value stored somewhere in your computer's memory. A **Pointer** is a special variable that "points" to the memory address of another variable. We use `&` to find the address and `*` to "dereference" (look inside) that address to see the value. This is useful for passing data around without copying it.

## Problem Description
Declare an integer and a pointer to that integer. Print the value of the integer, its memory address, and the value the pointer points to (dereferencing).
