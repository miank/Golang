// Write a Go program that fetches the lengths of multiple URLs concurrently using HTTP GET requests. The program should:

// Define a function FetchURL(url string, ch chan<- Result, wg *sync.WaitGroup) that sends a Result struct (containing the URL and its response body length or an error) to a channel.
// Use goroutines to fetch URLs concurrently.
// Handle errors for invalid URLs or failed requests.
// In main, process a slice of URLs, collect results via the channel, and print each URL’s response body length or error.
// Use sync.WaitGroup to ensure all goroutines complete before closing the channel.

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Result struct {
	URL    string
	Length int
	Err    error
}

func FetchURL(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching the output %v ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- Result{
			URL:    url,
			Length: 0,
			Err:    fmt.Errorf("non - 200 status code: %d", resp.StatusCode),
		}
		return
	}

	// reading the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{URL: url, Length: 0, Err: fmt.Errorf("reading body: %v", err)}
		return
	}

	ch <- Result{URL: url, Length: len(body), Err: nil}
}

func main() {

	urls := []string{
		"https://google.com",
		"https://apple.com",
		//"invalid.com",
	}

	ch := make(chan Result, len(urls))
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, v := range urls {
		go FetchURL(v, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Printf("URL: %s, Length: %d, Error: %v\n", result.URL, result.Length, result.Err)
	}
}
