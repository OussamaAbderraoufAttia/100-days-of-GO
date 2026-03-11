# Day 89 - Using an ORM (Intro to GORM)

## Goal
Simplify database work.

## Simple Explanation
Writing manually SQL can be tedious and error-prone. An **ORM** (Object-Relational Mapping) like **GORM** lets you treat database tables like regular Go structs. Instead of writing `INSERT INTO users...`, you just say `db.Create(&user)`. GORM handles all the SQL for you, which makes your code much cleaner and faster to write.

## Problem Description
Add the `gorm.io/gorm` package to your project. Define a struct and use GORM to automatically create the table ("AutoMigrate") and insert a record.
