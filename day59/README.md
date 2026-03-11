# Day 59 - Embedding Interfaces

## Goal
Compose behaviors.

## Simple Explanation
Just like you can put one struct inside another, you can put one interface inside another! This is called **Interface Embedding**. It lets you create complex interfaces by simply listing smaller ones. If a type satisfies all the small interfaces, it automatically satisfies the big one too.

## Problem Description
Define interfaces `Reader` and `Writer`. Then define a `ReadWriter` interface that embeds both. Implement a struct that satisfies `ReadWriter`.
