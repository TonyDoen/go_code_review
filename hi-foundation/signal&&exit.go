package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

// Sometimes we’d like our Go programs to intelligently handle Unix signals. For example, we might want a server to gracefully shutdown when it receives a SIGTERM, or a command-line tool to stop processing input if it receives a SIGINT.
func main() {
	sigs := make(chan os.Signal, 1)                      // Go signal notification works by sending os.Signal values on a channel. We’ll create a channel to receive these notifications (we’ll also make one to notify us when the program can exit).
	done := make(chan bool, 1)                           //
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // signal.Notify registers the given channel to receive notifications of the specified signals.
	go func() {                                          // This goroutine executes a blocking receive for signals. When it gets one it’ll print it out and then notify the program that it can finish.
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal") // The program will wait here until it gets the expected signal (as indicated by the goroutine above sending a value on done) and then exit.
	<-done
	fmt.Println("exiting")

	// Use os.Exit to immediately exit with a given status.
	defer fmt.Println("!") // defers will not be run when using os.Exit, so this fmt.Println will never be called.
	os.Exit(3)             // Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.
}

// Command
// $ go run signals. // When we run this program it will block waiting for a signal. By typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print interrupt and then exit.

// Command
// $ go run exit.go
// exit status 3
// $ go build exit.go // By building and executing a binary you can see the status in the terminal.
// $ ./exit
// $ echo $?
// 3
