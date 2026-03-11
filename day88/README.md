# Day 88 - CRUD Operations with a Database

## Goal
Manage database records.

## Simple Explanation
**CRUD** stands for **C**reate, **R**ead, **U**pdate, and **D**elete—the four basic things you do with any data. We use SQL commands like `INSERT`, `SELECT`, and `DELETE` through Go's `db.Exec()` or `db.Query()` functions. This is how you build a real application where user data is saved even if the server restarts.

## Problem Description
Using the database connection from Exercise 87, create a table named `users`. Write functions to insert a new user and retrieve all users from the table.
