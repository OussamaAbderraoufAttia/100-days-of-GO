# Day 78 - Using GOMAXPROCS

## Goal
Control the number of CPU cores used.

## Simple Explanation
Go usually handles core usage for you, but sometimes you want to limit how much of your computer it uses. `GOMAXPROCS` tells Go how many physical CPU cores it should allow your goroutines to run on at the same time. By default, it uses every core available, making Go incredibly fast on modern computers.

## Problem Description
Learn about `runtime.GOMAXPROCS`. Write a program that prints the current number of CPU cores being used by Go.
