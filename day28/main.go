package main

import "fmt"

type ContactInfo struct {
	Phone string
	City  string
}

type User struct {
	Name        string
	Age         int
	ContactInfo // Embedded struct
}

func main() {
	u := User{
		Name: "Alice",
		Age:  30,
		ContactInfo: ContactInfo{
			Phone: "555-0199",
			City:  "New York",
		},
	}

	// We can access fields of ContactInfo directly on the User!
	fmt.Printf("%s lives in %s. Contact: %s\n", u.Name, u.City, u.Phone)
}
