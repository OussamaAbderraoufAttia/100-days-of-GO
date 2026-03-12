# Day 82 - Routing with Search Parameters

## Goal
Read data from the URL query.

## Simple Explanation
When you see a URL like `?name=Alice`, the part after the `?` is called a **Query Parameter**. In Go, we can easily read these using `r.URL.Query().Get("name")`. This lets us create dynamic pages that respond to user input without needing complex forms.

## Problem Description
Modify your HTTP server to have a `/hello` path. Read a `name` parameter from the query string and respond with "Hello, [name]!". Default to "Guest" if no name is provided.
-
