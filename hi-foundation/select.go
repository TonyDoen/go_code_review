package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
func main() {
	// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select { // We’ll use select to await both of these values
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}

// > Output:
// command-line-arguments
// received one
// received two
// > Elapsed: 2.296s  //Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
// > Result: Success
