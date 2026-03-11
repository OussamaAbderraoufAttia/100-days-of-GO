# Day 100 - Final Capstone Project: URL Shortener

## Goal
Build a complete real-world application.

## Simple Explanation
Congratulations! For your final project, you'll build a **URL Shortener** (like Bitly).
It needs:
1. A path to take a long URL and return a short code.
2. A database to save where each short code points to.
3. A path that takes a short code and "redirects" the user to the real website.
This combines everything: HTTP, JSON, Databases, and Logic!

## Problem Description
Build a complete URL shortener. 
- `POST /shorten`: Takes a long URL, generates a random short code, and saves it.
- `GET /[short_code]`: Redirects the browser to the original URL.
- Use an SQLite database to store the mappings.
