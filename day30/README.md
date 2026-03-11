# Day 30 - Pointer Receivers

## Goal
Modify struct state in methods.

## Simple Explanation
By default, if you pass a struct to a method, Go makes a **copy** of it. If you change something inside the method, the original stays the same. To change the *original* struct, we use a **Pointer Receiver**. We write it with an asterisk like `func (p *Person) Change()`. This tells Go to point to the real thing in memory instead of making a copy.

## Problem Description
Add a method `HaveBirthday()` to the `Person` struct that increments their age by 1. Use a pointer receiver so the change persists.
