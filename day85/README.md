# Day 85 - Basic REST API (GET)

## Goal
Build a data endpoint.

## Simple Explanation
A **REST API** is a way for computers to share data with each other. Instead of returning a webpage for humans, we return **JSON** for other programs. In this project, we create a list of items (like a list of movies) and write a handler that turns that list into a JSON string when someone visits `/api/movies`.

## Problem Description
Define a `Movie` struct with a few fields. Create a slice of movies. Write a handler `/movies` that returns this slice as JSON data.
