package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "user_data.db")
	defer db.Close()

	// 1. CREATE table
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")

	// 2. INSERT (Create)
	db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")

	// 3. QUERY (Read)
	rows, _ := db.Query("SELECT id, name FROM users")
	defer rows.Close()

	fmt.Println("User List:")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("[%d] %s\n", id, name)
	}
}
