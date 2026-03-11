package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my-secret-key")

func createToken(username string) (string, error) {
	// Create claims (the data inside the token)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expires in 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func main() {
	token, _ := createToken("Alice")
	fmt.Println("Generated Token:", token)
}
