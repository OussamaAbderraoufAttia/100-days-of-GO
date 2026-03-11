package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mySuperSecretPassword123"

	// 1. Hash the password
	// Cost is the 'strength' of the encryption (higher is slower/safer)
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println("Example Hash:", string(hash))

	// 2. Compare a plain-text password with the hash
	testPassword := "mySuperSecretPassword123"
	err := bcrypt.CompareHashAndPassword(hash, []byte(testPassword))

	if err == nil {
		fmt.Println("Success: Password matches!")
	} else {
		fmt.Println("Error: Invalid password.")
	}
}
