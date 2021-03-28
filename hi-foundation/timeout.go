package main

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.
func main() {
	//timeout 1 scenario
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2) // returns its result on a channel ch1 after 2s
		ch1 <- "result 1"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second * 1): // <-Time.After awaits a value to be sent after the timeout of 1s.
		fmt.Println("timeout 1")
	}

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second * 3): // <-Time.After awaits a value to be sent after the timeout of 3s. we allow a longer timeout of 3s, then the receive from ch1 will succeed and we’ll print the result.
		fmt.Println("timeout 1")
	}
}
