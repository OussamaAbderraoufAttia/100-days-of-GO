# Day 40 - JSON Deserialization (Unmarshal)

## Goal
Convert JSON back to structs.

## Simple Explanation
Receiving data from the web often means getting a JSON string. We use `json.Unmarshal` to turn that string back into a Go struct. We must provide a pointer to the struct so Go can fill it with data. This is the opposite of Marshal.

## Problem Description
Take a JSON string representing a User and unmarshal it into a `User` struct.
-
