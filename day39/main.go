package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func main() {
	p := Product{Name: "Laptop", Price: 999.99, Stock: 10}

	// Convert to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
