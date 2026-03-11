# Day 77 - Profiling Basics (pprof)

## Goal
Find where your code is slow.

## Simple Explanation
**Profiling** is like taking an X-ray of your program to see which functions are taking the most time or using the most memory. Go has a tool called `pprof` that generates reports showing you exactly where the "bottlenecks" (slow parts) are. This lets you fix only the parts that actually matter for speed.

## Problem Description
Research how to use `net/http/pprof` to profile a Go application. Write a basic setup for exposing pprof via an HTTP server.
