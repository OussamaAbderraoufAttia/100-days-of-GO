# Day 87 - Connecting to a Database (SQL basics)

## Goal
Persistent data storage.

## Simple Explanation
Variables in Go disappear when you stop the program. To save data forever, we use a **Database**. Go's `database/sql` package is a standard way to talk to many different types of databases (like SQLite or Postgres). You need a "Driver" for your specific database, and then you use `sql.Open` to create a connection.

## Problem Description
Research how to use the `database/sql` package with an SQLite driver (e.g., `modernc.org/sqlite`). Write a program that opens a connection to a local database file named `test.db`.
-
