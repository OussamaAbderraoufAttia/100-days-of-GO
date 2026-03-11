package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://google.com/search?q=golang&lang=en"

	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host  :", u.Host)
	fmt.Println("Path  :", u.Path)
	
	// Query parameters
	queryParams := u.Query()
	fmt.Println("Query Parameter 'q':", queryParams.Get("q"))
	fmt.Println("Query Parameter 'lang':", queryParams.Get("lang"))
}
