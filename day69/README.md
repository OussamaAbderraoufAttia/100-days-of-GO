# Day 69 - Context Basics (Timeout)

## Goal
Manage lifecycles and timeouts.

## Simple Explanation
In real-world apps, you shouldn't wait forever for something (like a website to load). The **Context** package is used to pass "deadlines" and "signals" through your program. If a task takes too long, the Context will trigger a "done" signal, telling all workers to stop and give up, which saves resources.

## Problem Description
Create a context with a 2-second timeout. Use a `select` statement to simulate a task that takes 3 seconds and Observe the timeout triggering.
