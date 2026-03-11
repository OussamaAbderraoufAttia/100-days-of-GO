# Day 55 - Working with Time

## Goal
Handle dates and times.

## Simple Explanation
Go's `time` package is used for everything related to dates and times. A funny thing in Go is how you format time: you use a specific "benchmark date" (January 2, 2006, at 3:04:05 PM) to show Go the exact pattern you want. `time.Now()` gives you the current moment.

## Problem Description
Write a program that prints the current date and time in a custom format (e.g., "YYYY-MM-DD HH:MM:SS").
