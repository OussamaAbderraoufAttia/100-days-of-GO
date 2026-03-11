package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// 1. Define subcommands
	greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
	nameFlag := greetCmd.String("name", "World", "name to greet")

	timeCmd := flag.NewFlagSet("time", flag.ExitOnError)

	// 2. Check which subcommand was used
	if len(os.Args) < 2 {
		fmt.Println("Expected 'greet' or 'time' subcommands")
		os.Exit(1)
	}

	// 3. Parse the correct subcommand
	switch os.Args[1] {
	case "greet":
		greetCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *nameFlag)
	case "time":
		timeCmd.Parse(os.Args[2:])
		fmt.Println("Current time is:", time.Now().Format("15:04:05"))
	default:
		fmt.Println("Unknown subcommand")
		os.Exit(1)
	}
}
