# Day 23 - Slice Length and Capacity

## Goal
Understand how slices grow.

## Simple Explanation
Every slice has a **Length** (how many items it currently has) and a **Capacity** (how much space it has in the underlying memory). When you `append` to a slice that is "full" (Length equals Capacity), Go automatically creates a new, bigger slice and copies your data over. We can use the `make()` function to set an initial size and capacity.

## Problem Description
Create a slice using `make` with a length of 0 and a capacity of 5. In a loop, append 10 integers to it. In each iteration, print the element, the current length, and the current capacity.
