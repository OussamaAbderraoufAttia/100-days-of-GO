# Day 43 - Type Assertions

## Goal
Extract underlying values from interfaces.

## Simple Explanation
When you have a variable of type `any`, Go doesn't know what's inside it until you ask. **Type Assertion** is how you say: "I think this interface currently holds a string, let me try to get it out." We use a special syntax `v, ok := i.(string)`. The `ok` variable tells us if our guess was correct, so our program doesn't crash if we are wrong.

## Problem Description
Create an interface variable holding a string. Use a type assertion to extract the string value. Handle the case where the assertion might fail using the "comma ok" idiom.
