# Day 71 - Worker Pools

## Goal
Manage multiple concurrent jobs efficiently.

## Simple Explanation
A **Worker Pool** is like a team of workers waiting for tasks. You create a fixed number of goroutines (the "workers") that all listen to the same "jobs" channel. This prevents your program from starting too many goroutines at once and crashing. It's a great way to process a huge list of tasks without overwhelming your computer.

## Problem Description
Implement a worker pool with 3 workers. Send 10 jobs (represented by integers) into a channel and have the workers process them concurrently.
