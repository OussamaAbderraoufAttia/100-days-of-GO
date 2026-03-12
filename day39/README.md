# Day 39 - JSON Serialization (Struct Tags)

## Goal
Convert structs to JSON.

## Simple Explanation
JSON is the standard language of the web. Go can easily convert structs into JSON using "Struct Tags". These are little notes in backticks next to field names that tell Go what the JSON keys should look like (usually lowercase). We use `json.Marshal` to do the conversion.

## Problem Description
Define a `Product` struct with fields for Name, Price, and Stock. Use struct tags to make the JSON keys lowercase. Marshal a `Product` instance into a JSON string and print it.
-
