# Day 91 - JWT Authentication Basics

## Goal
Secure your API with tokens.

## Simple Explanation
**JWT** (JSON Web Token) is a tiny piece of data that proves who someone is. When a user logs in, the server gives them a JWT. For every request after that, the user "shows" their token. This is much better than passwords because the token is temporary and the server doesn't have to look up the user's password in the database every single time they want to see a page.

## Problem Description
Research the `golang-jwt/jwt` package. Write a program that creates a digital "token" with a username inside it, signs it with a secret key, and then prints the token.
