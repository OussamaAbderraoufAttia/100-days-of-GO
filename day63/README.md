# Day 63 - Channels Basics

## Goal
Communicate between goroutines.

## Simple Explanation
If Goroutines are workers, **Channels** are the pipes they use to talk to each other. One goroutine can "send" a value into a pipe using `channel <- value`, and another can "receive" it using `value := <- channel`. This is the safest way to shared data because only one goroutine has the data at a time.

## Problem Description
Create a channel of type `string`. Start a goroutine that sends "Hello from Channel" into it. In the main function, receive and print the message.
