# Day 92 - Using Docker with Go

## Goal
Package your app for any computer.

## Simple Explanation
**Docker** is like a shipping container for your code. It packages your app and everything it needs (like specific libraries) into one "Image". This means your app will run exactly the same way on your laptop, your friend's laptop, and a giant server in the cloud, without you ever hearing the phrase "but it worked on my machine!"

## Problem Description
Write a basic `Dockerfile` for a Go application. Use a "multi-stage build" to keep the final image size small.
