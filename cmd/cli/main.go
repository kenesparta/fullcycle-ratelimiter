package main

import (
	"flag"
	"fmt"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var (
		url         = flag.String("url", "", "URL to test")
		maxRequests = flag.Int("r", 100, "Maximum amount of requests to send")
		timeWindow  = flag.Int("t", 1, "Time in seconds of each request")
		method      = flag.String("m", "GET", "HTTP method to use")
		apiKey      = flag.String("k", "", "API Key for the request")
	)
	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		return
	}

	performRequestsAtRate(*url, *method, *apiKey, (*maxRequests)/(*timeWindow))
}

func performRequestsAtRate(url, method, apiKey string, ratePerSecond int) {
	ticker := time.NewTicker(time.Second / time.Duration(ratePerSecond))
	defer ticker.Stop()

	var wg sync.WaitGroup
	for {
		select {
		case <-ticker.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				sendRequest(url, method, apiKey)
			}()
		}
	}
}

func sendRequest(url, method, apiKey string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("Error creating request: %s\n", err)
		return
	}

	if apiKey != "" {
		req.Header.Add(entity.APIKeyHeaderName, apiKey)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Request error: %s\n", err)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		return
	}

	log.Printf("Status Code: %d\n", response.StatusCode)
	log.Printf("Body: %s\n", body)
	defer response.Body.Close()
}
