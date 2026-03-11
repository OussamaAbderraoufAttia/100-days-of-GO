package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	// errgroup helps manage a group of goroutines that might fail
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("Task 1 running...")
		return nil
	})

	g.Go(func() error {
		fmt.Println("Task 2 failing...")
		return errors.New("boom! task 2 crashed")
	})

	g.Go(func() error {
		fmt.Println("Task 3 running...")
		return nil
	})

	// Wait for all tasks and catch the first error
	if err := g.Wait(); err != nil {
		fmt.Printf("Finished with error: %v\n", err)
	} else {
		fmt.Println("Successfully finished all tasks.")
	}
}
