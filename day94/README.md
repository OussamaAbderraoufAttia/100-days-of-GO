# Day 94 - Graceful Shutdown

## Goal
Stop your server politely.

## Simple Explanation
When you stop a server (like pressing Ctrl+C), you shouldn't just cut the power immediately. Some users might be in the middle of a purchase or a download! **Graceful Shutdown** tells the server: "Stop accepting NEW users, but wait for the CURRENT users to finish their tasks before you truly turn off." This prevents data loss and angry customers.

## Problem Description
Modify your HTTP server to listen for "interrupt signals" (like Ctrl+C). Use `srv.Shutdown()` with a context to wait for active requests to finish before exiting.
-
