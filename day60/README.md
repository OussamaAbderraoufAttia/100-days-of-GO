# Day 60 - Simple File Metadata Tool

## Goal
Practical application of file and CLI concepts.

## Simple Explanation
This project combines your knowledge of the file system and command-line tools. We use `os.Stat` to get "metadata" about a file—things like how big it is, when it was last changed, and who is allowed to read it. This is exactly how tools like `ls` on Linux or `Dir` on Windows work.

## Problem Description
Build a CLI tool that takes a file path as an argument and prints its size, last modified time, and permissions. Use `os.Stat`.
-
