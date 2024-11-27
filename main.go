package main

import (
	"cocurrency/fetcher"
	"fmt"
	"sync"
)

func main() {
	// List of API endpoints to fetch data from
	apis := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
	}

	// Create a WaitGroup to wait for all GoRoutines to finish
	var wg sync.WaitGroup
	// Create a channel to receive API results
	ch := make(chan string)

	// Mutex to control access to shared resources (e.g., logging)
	var mu sync.Mutex

	// Start a GoRoutine for each API call
	for _, api := range apis {
		wg.Add(1)
		go fetcher.FetchAPI(api, ch, &wg, &mu)
	}

	// Close the channel once all GoRoutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Wait for results from the channel
	for result := range ch {
		// Log the result in a thread-safe way using Mutex
		mu.Lock()
		fmt.Println(result)
		mu.Unlock()
	}

	fmt.Println("All workers are done!")
}
