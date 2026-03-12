# Day 95 - Environment-based Config (Cleanenv)

## Goal
Manage production-ready configuration.

## Simple Explanation
In a big app, you have dozens of settings (database URLs, API keys, etc.). The `cleanenv` library helps you manage these by letting you define a struct for your config. It automatically looks for these settings in environment variables or a `.env` file, which makes it very easy to change how your app behaves on your laptop versus a production server.

## Problem Description
Research the `github.com/ilyakaznacheev/cleanenv` package. Create a configuration struct and load values from environment variables.
-
