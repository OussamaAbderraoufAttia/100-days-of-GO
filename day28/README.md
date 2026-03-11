# Day 28 - Embedded Structs (Composition)

## Goal
Use composition instead of inheritance.

## Simple Explanation
In many languages, objects "inherit" from each other. In Go, we use "Composition" by putting one struct inside another. This is called **Embedding**. If you embed a struct, the outer struct can use all the fields of the inner one as if they were its own. This is a very clean way to share behavior.

## Problem Description
Define a `ContactInfo` struct. Then, embed `ContactInfo` inside a `User` struct. Create a `User` and access the embedded fields directly.
