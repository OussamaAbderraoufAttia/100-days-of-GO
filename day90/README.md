# Day 90 - Password Hashing (bcrypt)

## Goal
Secure user credentials.

## Simple Explanation
Never save a user's real password in a database! If a hacker steals your data, they will have everyone's passwords. Instead, we use **Hashing** (via the `bcrypt` library). This turns a password into a scrambled string that cannot be turned back. To check a password, you hash what the user typed and compare it to the scrambled string in your database.

## Problem Description
Use the `golang.org/x/crypto/bcrypt` package. Write a program that hashes a password and then verifies if a user-provided string matches that hash.
-
