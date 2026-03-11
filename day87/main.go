package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Blank import for the driver
)

func main() {
	// Open connection to test.db (creates it if it doesn't exist)
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	fmt.Println("Successfully connected to the database!")
}
