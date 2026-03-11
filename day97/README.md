# Day 97 - Building a Simple CRUD API (Part 1)

## Goal
Architecture and Project Layout.

## Simple Explanation
When building a real app, don't just put everything in one file. Follow the "Standard Go Project Layout". Group your logic into folders like `cmd/` for starting the app, `internal/` for your secret recipes, and `pkg/` for tools you want to share. This makes your code easy to navigate for other developers.

## Problem Description
Set up a project folder with `cmd`, `internal/handlers`, and `internal/models`. Define a model for a "Book" (ID, Title, Author) in the models folder.
