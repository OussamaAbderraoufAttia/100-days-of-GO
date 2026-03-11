# Day 56 - Environment Variables

## Goal
Configure apps using the environment.

## Simple Explanation
Instead of writing secrets (like passwords or port numbers) directly in your code, we use **Environment Variables**. These are settings on the computer where the code is running. Use `os.Getenv` to read them. This makes your apps much safer and easier to move between different servers.

## Problem Description
Write a program that reads an environment variable `APP_PORT`. If it's not set, use a default value of `8080`.
