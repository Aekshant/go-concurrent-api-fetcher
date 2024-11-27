package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// FetchAPI is the function that fetches data from a given API
func FetchAPI(apiURL string, ch chan<- string, wg *sync.WaitGroup, mu *sync.Mutex) {
	// Defer the Done method to signal that the GoRoutine is finished
	defer wg.Done()

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		ch <- fmt.Sprintf("API %s: Error fetching data - %s", apiURL, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body using io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("API %s: Error reading response body - %s", apiURL, err)
		return
	}

	// Return the result through the channel
	ch <- fmt.Sprintf("API %s: Fetched data successfully, Response: %s", apiURL, body)
}
