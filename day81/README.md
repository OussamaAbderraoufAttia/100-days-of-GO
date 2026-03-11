# Day 81 - Basic HTTP Server (net/http)

## Goal
Create your first web server.

## Simple Explanation
Go was built for the internet, and its `net/http` package is incredibly powerful. To start a server, you only need to define a **Handler** (a function that decides what to send back to the user) and then tell Go to "Listen and Serve" on a specific port (like 8080). This is the foundation of almost every web app or API written in Go.

## Problem Description
Write a program that starts an HTTP server on port 8080. It should have one path `/` that returns "Welcome to my Go Web Server!" as a text response.
