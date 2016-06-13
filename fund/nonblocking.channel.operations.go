package main

import "fmt"

func main() {
	message := make(chan string)
	signal := make(chan bool)

	// go func() { message <- "false" }()

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
