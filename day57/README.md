# Day 57 - Working with URLs

## Goal
Parse and build URLs.

## Simple Explanation
The web runs on URLs. Go's `net/url` package helps you take a URL string and break it into pieces, like the "Scheme" (http), the "Host" (google.com), and the "Query" (things after the ?). This is essential for building web scrapers or API clients.

## Problem Description
Take a URL string (e.g., "https://example.com/search?q=golang") and parse it using the `net/url` package. Print the scheme, host, path, and query parameters.
