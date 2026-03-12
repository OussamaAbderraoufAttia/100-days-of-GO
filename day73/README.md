# Day 73 - Tickers and Timers

## Goal
Schedule recurring or delayed events.

## Simple Explanation
- A **Timer** is for doing something *once* in the future (like an alarm clock).
- A **Ticker** is for doing something *repeatedly* (like the ticking of a clock).
Both use channels to send you a message when it's time to act. Don't forget to `Stop()` them when you are done to save resources!

## Problem Description
Create a ticker that prints "Tick" every 500ms and a timer that stops the program after 2 seconds.
-
