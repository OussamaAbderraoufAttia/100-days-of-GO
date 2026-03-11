package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // The underscore import registers the pprof handlers automatically
)

func main() {
	fmt.Println("Pprof server starting on :6060")
	fmt.Println("Visit http://localhost:6060/debug/pprof/ in your browser")

	// Start a separate web server just for profiling info
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Simulate your main application running
	for {
		// Do some "work" here
	}
}
