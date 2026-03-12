# Day 72 - Rate Limiting

## Goal
Control the speed of operations.

## Simple Explanation
**Rate Limiting** is about making sure your program doesn't do too many things too quickly (like sending 1000 emails in one second). In Go, we use tickers or channels to create a "clock" that only allows an operation to happen once every X milliseconds. This is vital for playing nice with other services and APIs.

## Problem Description
Implement a simple rate limiter that only allows 5 tasks to be processed per second.
-
