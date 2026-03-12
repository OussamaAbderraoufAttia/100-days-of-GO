# Day 64 - Buffered Channels

## Goal
Handle asynchronous communication.

## Simple Explanation
Normal channels are "unbuffered," meaning the sender has to wait for a receiver to be ready. **Buffered Channels** have a "mailbox" size. You can send a few items into the pipe even if no one is there to grab them yet. The sender only has to wait if the mailbox is completely full.

## Problem Description
Create a buffered channel with a capacity of 2. Send two values into it without a receiver, then receive and print them.
-
