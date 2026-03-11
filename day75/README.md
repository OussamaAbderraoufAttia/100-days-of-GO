# Day 75 - Sync.Pool

## Goal
Reuse objects to save memory.

## Simple Explanation
Creating and destroying objects many times can be hard on your computer's memory (the "Garbage Collector"). `sync.Pool` is like a "lost and found" box for objects. Instead of throwing away an object when you are done, you "Put" it in the pool. When you need a new one, you "Get" it from the pool. If the pool is empty, it creates a new one. This saves a lot of time and memory in high-performance apps.

## Problem Description
Use `sync.Pool` to reuse a slice of bytes across multiple operations.
