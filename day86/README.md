# Day 86 - REST API (POST/JSON input)

## Goal
Receive data via API.

## Simple Explanation
To add new data to our API, we use the **POST** method. The user sends a JSON object in the "body" of their request. We use `json.NewDecoder` to read that JSON and turn it into a Go struct. This is how you "upload" information, like creating a new user or posting a comment.

## Problem Description
Create a handler that accepts a `POST` request to `/movies`. Parse the incoming JSON to add a new movie to your in-memory slice.
