package main

import (
	"fmt"
	"time"
)

//
func worker(done chan bool) { // The done channel will be used to notify another goroutine that this function’s work is done.
	fmt.Println("worker is working now.")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // Send a value to notify that we’re done.
}

func ping(sender chan<- string, msg string) { // you can specify if a channel is meant to only send or receive values. This ping function only accepts a channel for sending values.
	sender <- msg
}

func pong(receiver <-chan string, sender chan<- string) { // The pong function accepts one channel for receives (ps) and a second for sends (pos).
	msg := <-receiver
	sender <- msg
	// sender <- <-receiver //funk but bingo
}

func main() {
	//use channel
	message := make(chan string)              // Create a new channel with make(chan val-type).
	go func() { message <- "pingpong" }()     // Send a value into a channel using the channel <- syntax. Here we send "pingpong" to the messages channel from a new goroutine.
	msg := <-message                          // The <-channel syntax receives a value from the channel.
	fmt.Println(msg, &msg, message, &message) // Here we’ll receive the "pingpong" message we sent above and print it out.

	//channel buffering
	message = make(chan string, 3) // Here we make a channel of strings buffering up to 3 values.
	message <- "buffered"          // we can send these values into the channel without a corresponding concurrent receive.
	message <- "channel"
	message <- "haha"
	fmt.Println(<-message, message, &message) // we can receive these values as usual
	fmt.Println(<-message, message, &message)
	fmt.Println(<-message, message, &message)

	message <- "this" // we can send these values into the channel without a corresponding concurrent receive.
	message <- "is"
	message <- "test"
	fmt.Println(<-message, message, &message)
	fmt.Println(<-message, message, &message)
	fmt.Println(<-message, message, &message)

	//channel synchronization //同步
	done := make(chan bool, 1)       //
	go worker(done)                  // Start a worker goroutine, giving it the channel to notify on.
	fmt.Println(<-done, done, &done) // "<-done" block until we receive a notification from the worker on the channel. If you removed the <- done line from this program, the program would exit before the worker even started.

	//channel direction
	pingt := make(chan string, 1)
	pongt := make(chan string, 1)
	ping(pingt, "ping pass message")
	pong(pingt, pongt)
	fmt.Println(<-pongt, pongt, &pongt)
}
