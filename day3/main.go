package main

import "fmt"

const (
	// iota starts at 0 and increments for each line
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	const SpeedOfLight = 299792458 // meters per second

	fmt.Println("Days of the week values:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

	fmt.Printf("The speed of light is %d m/s\n", SpeedOfLight)
}
