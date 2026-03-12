# Day 83 - Handlers and Middleware

## Goal
Process requests before they reach your main logic.

## Simple Explanation
**Middleware** is code that runs before your main logic (like a security guard at a door). It can check if a user is logged in, log every request to the console, or add special headers to a response. In Go, middleware is just a function that wraps another handler.

## Problem Description
Create a logging middleware that prints the requested URL path to the console for every single request before calling the actual handler.
-
