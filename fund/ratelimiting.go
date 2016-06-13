package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.
func main() {
	// Suppose we want to limit our handling of incoming requests. We’ll serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200) // This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.

	for req := range requests {
		<-limiter // By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("bursty...")
	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.
	burstylimiter := make(chan time.Time, 3) // We can accomplish this by buffering our limiter channel. This burstylimiter channel will allow bursts of up to 3 events.
	for i := 0; i < 3; i++ {                 // Fill up the channel to represent allowed bursting.
		burstylimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) { // Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
			burstylimiter <- t
		}
	}()

	burstyrequests := make(chan int, 50)
	for i := 1; i <= 50; i++ {
		burstyrequests <- i
	}
	close(burstyrequests)
	for req := range burstyrequests {
		<-burstylimiter
		fmt.Println("request", req, time.Now())
	}

	// Running our program we see the first batch of requests handled once every ~200 milliseconds as desired.
	// For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each
}
