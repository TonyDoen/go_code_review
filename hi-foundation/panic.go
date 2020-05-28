package main

import "os"

// A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
func main() {
	// panic("a problem.")
	_, err := os.Create("/tmp/tmp.file.0")
	if err != nil { // A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
		panic(err)
	}
}

// Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
