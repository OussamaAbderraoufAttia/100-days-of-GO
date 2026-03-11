package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model // Adds ID, CreatedAt, etc.
	Code  string
	Price uint
}

func main() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Migrate the schema (creates the table)
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key 1

	fmt.Printf("Product Found: %s, Price: %d\n", product.Code, product.Price)
}
