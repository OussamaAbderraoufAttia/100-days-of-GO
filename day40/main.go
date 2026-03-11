package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// JSON data as a string
	jsonInput := `{"id": 1, "name": "John Doe", "email": "john@example.com"}`

	var u User
	// Convert string to bytes and unmarshal
	err := json.Unmarshal([]byte(jsonInput), &u)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("User ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
}
