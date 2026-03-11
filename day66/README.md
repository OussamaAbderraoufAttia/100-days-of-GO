# Day 66 - Select Statement (Multiplexing)

## Goal
Wait on multiple channels.

## Simple Explanation
The `select` statement is like a `switch`, but for channels. It lets a goroutine wait on many pipes at the same time. Whichever pipe gets a message first, that `case` will run. This is extremely useful for things like timeouts or listening to multiple data sources simultaneously.

## Problem Description
Create two channels that receive messages at different times. Use a `select` statement to print whichever message arrives first.
