# Day 41 - Interfaces Basics

## Goal
Define behavior through interfaces.

## Simple Explanation
An **Interface** in Go is a contract or a list of specific behaviors (methods). It doesn't tell a type *how* to do something, just that it *must* have those methods. If a struct has all the methods listed in an interface, Go automatically considers that struct to "implement" the interface. This lets us write flexible code that works with many different types as long as they "interface" the same way.

## Problem Description
Define a `Shape` interface with an `Area() float64` method. Create `Circle` and `Rectangle` structs that implement this interface.
-
