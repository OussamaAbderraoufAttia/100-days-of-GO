# Day 96 - Logging with Zerolog or Zap

## Goal
Produce high-performance logs.

## Simple Explanation
Standard Go logs are fine for small projects, but for big apps, you want **Structured Logging**. Libraries like `zerolog` produce logs in JSON format. This is amazing because it's easy for robots and search tools (like ElasticSearch) to read your logs and find exactly where a specific error occurred in seconds.

## Problem Description
Research the `github.com/rs/zerolog` package. Write a program that logs a message with extra fields (like "user_id") in JSON format.
-
