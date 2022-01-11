package main

import (
	"fmt"
	"sync"
)

// This one doesn't know anything about goroutines
func runMe() {
	fmt.Println("Hello from goroutine")
}

// Good idea to seaprate concurrent code from business logic
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		runMe()
		wg.Done()
	}()

	wg.Wait()
}
