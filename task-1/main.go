package main

import (
	"fmt"
	"net/http"
	"sync"
)

// We use a Mutex to safely update this shared counter
var (
	count      int
	countMutex sync.Mutex
)

func fetchPost(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %d", id)
		return
	}
	defer resp.Body.Close()

	// LOCK: Only one goroutine can increment the counter at a time
	countMutex.Lock()
	count++
	countMutex.Unlock()

	ch <- fmt.Sprintf("Post %d: Status %s", id, resp.Status)
}

func main() {
	postsToFetch := 5
	results := make(chan string, postsToFetch)
	var wg sync.WaitGroup

	for i := 1; i <= postsToFetch; i++ {
		wg.Add(1)
		// Launching a Goroutine
		go fetchPost(i, results, &wg)
	}

	// Close channel once all Goroutines are finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Receiving from the Channel
	for res := range results {
		fmt.Println(res)
	}

	fmt.Printf("Total successful increments via Mutex: %d\n", count)
}
