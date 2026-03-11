# Day 80 - Building a Concurrent Link Checker

## Goal
Capstone concurrency project.

## Simple Explanation
In this project, we put everything together to build something useful: a tool that checks many websites at once to see if they are "up" or "down". Instead of checking one by one (which is slow), we use Goroutines for speed and Channels to send the results back to the main program. This is a classic "real-world" Go project.

## Problem Description
Write a program that takes a slice of 5-10 URLs. For each URL, start a goroutine that performs an `http.Get`. Use a channel to send back a message stating whether the site is up (status code 200) or down.
