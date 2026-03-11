# Day 24 - Iterating over Slices (Range)

## Goal
Use range for cleaner loops.

## Simple Explanation
Instead of using a traditional `for` loop with an index like `i++`, Go provides the `range` keyword. It makes it easy to go through every item in a slice or map. `range` gives you two things each time: the current **index** and the **value** at that index. If you don't need the index, you can use an underscore `_` to ignore it.

## Problem Description
Create a slice of integers. Use `for range` to iterate over the slice and print both the index and the value of each element.
