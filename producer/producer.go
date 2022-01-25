package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	do()
}

// sends concurrent requests to different enqueuing services
func do() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sendConcurrentRequests(100, 5000+i)
		}(i)
	}
	wg.Wait()
}

// sends n concurrent producer requests to the enqueuing service running at port "port"
func sendConcurrentRequests(n int, port int) {
	fmt.Printf("Sending %d concurrent requests on port %d\n", n, port)
	url := fmt.Sprintf("http://localhost:%d", port)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			body := fmt.Sprintf("\"happay\": \"%d\"", i)
			http.Post(url, "application/json", bytes.NewBuffer([]byte(body)))
		}(i + 1)
	}
	wg.Wait()
}
