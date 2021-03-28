package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Command
// $ echo 'hello'   > /tmp/lines
// $ echo 'filter' >> /tmp/lines
// $ cat /tmp/lines | go run linefilter.go
func main() {
	scanner := bufio.NewScanner(os.Stdin) // Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token; which is the next line in the default scanner.
	for scanner.Scan() {                  // Text returns the current token, here the next line, from the input.
		ucl := strings.ToUpper(scanner.Text()) //
		fmt.Println(ucl)                       // Write out the uppercased line.
	}

	if err := scanner.Err(); err != nil { // Check for errors during Scan. End of file is expected and not reported by Scan as an error.
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
