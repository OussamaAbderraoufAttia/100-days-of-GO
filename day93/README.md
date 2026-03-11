# Day 93 - Writing a Makefile

## Goal
Automate common tasks.

## Simple Explanation
A **Makefile** is a simple list of shortcuts for terminal commands. Instead of typing `go build -o myapp main.go` every time, you can just type `make build`. This saves time and ensures everyone on your team runs the exact same commands for building, testing, and cleaning the project.

## Problem Description
Create a `Makefile` with targets for `build` (compiles the app), `run` (compiles and starts), and `clean` (deletes the binary).
