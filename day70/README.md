# Day 70 - Context with Cancel

## Goal
Manually stop operations.

## Simple Explanation
Sometimes you want to stop a task yourself, not just wait for a timer. `context.WithCancel` gives you a `cancel` function. When you call it, it sends a message to all workers using that context, telling them to stop whatever they are doing immediately. This is how you gracefully shut down a server.

## Problem Description
Start three worker goroutines that listen to a `Done()` signal from a context. In the main function, call the `cancel()` function after 1 second and verify all workers stopped.
