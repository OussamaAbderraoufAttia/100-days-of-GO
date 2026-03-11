# Day 76 - Race Condition Detection

## Goal
Find concurrency bugs.

## Simple Explanation
A **Race Condition** is one of the hardest bugs to find because it only happens sometimes—when two workers touch the same data at the exact same millisecond. Go has a built-in "Race Detector". You run your tests with `go test -race`, and Go will watch your program and shout if it sees two workers fighting over the same variable!

## Problem Description
Write a program with a deliberate race condition (two goroutines incrementing a shared variable without a mutex). Learn how to use the `-race` flag to find it. (Note: Since you can't run code, just write the racey code).
