# Day 84 - Serving Static Files

## Goal
Deliver HTML, CSS, and images.

## Simple Explanation
Most websites aren't just text; they have images and stylesheets. Go's `http.FileServer` is a built-in tool that can automatically serve any file in a folder to the internet. Just tell it which directory to look at, and it handles all the complicated parts of reading files and sending them to the browser.

## Problem Description
Create a folder named `public` with an `index.html` file. Use `http.FileServer` to serve the contents of that folder on the root path `/`.
-
