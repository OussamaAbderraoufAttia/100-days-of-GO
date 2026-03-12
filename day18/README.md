# Day 18 - Scoping

## Goal
Understand variable scope.

## Simple Explanation
"Scope" is where a variable lives and can be seen. If you create a variable inside a function, it's a "local" variable and cannot be seen outside. If you create it outside of all functions, it's a "package-level" variable. If a local variable has the same name as a package variable, it "shadows" it inside that function.

## Problem Description
Write a program with a package-level variable and a function-level variable with the same name. Observe which one is accessed inside the function.
-
