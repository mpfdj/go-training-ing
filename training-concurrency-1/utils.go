package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"time"
)

var counter atomic.Int64

// Simulate a CPU intensive operation
var computeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for time.Since(start) <= time.Millisecond*500 {
		for i := 0; i < 1000; i++ {
			_ = i * i
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("operation finished"))
})

// Perform an HTTP request to the given endpoint
func doRequest(url string) error {
	client := &http.Client{Transport: http.DefaultTransport}

	response, err := client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	counter.Add(1)

	return nil
}

func startServer() (string, func()) {
	// Start a built-in webserver
	server := httptest.NewServer(computeHandler)

	// Start a new timer
	start := time.Now()

	return server.URL, func() {
		defer server.Close()
		// Print the time it took to perform the requests
		diff := time.Since(start)
		fmt.Printf("took %fs\n", diff.Seconds())

		// Verify that 100 requests were performed
		if counter.Load() != 100 {
			fmt.Fprintf(os.Stderr, "expected 1000 requests, got %d\n", counter.Load())
		}
	}
}
