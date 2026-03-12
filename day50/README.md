# Day 50 - Subcommands in CLI

## Goal
Build complex CLI structures.

## Simple Explanation
Advanced tools often have "subcommands" (like `git add` or `docker build`). In Go, we use `flag.NewFlagSet` to create a separate group of flags for each subcommand. This lets our program do many different things, each with its own specific options, all from the same single executable.

## Problem Description
Use the `flag` package to create a tool with two subcommands: `greet` (takes a name) and `time` (prints current time).
-
