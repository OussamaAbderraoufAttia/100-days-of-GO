# Day 65 - Closing Channels and Range

## Goal
Signal the end of a data stream.

## Simple Explanation
When a goroutine has finished sending all its data, it can `close()` the channel. This tells everyone else that no more messages are coming. You can use a `for range` loop on a channel to automatically keep receiving data until the channel is closed.

## Problem Description
Start a goroutine that sends numbers 1 to 5 into a channel and then closes it. Use a `for range` loop in the main function to print all values.
