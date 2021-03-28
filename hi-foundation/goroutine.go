package main

import "fmt"

// Command
// # go run goroutine.go

func ft1(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	ft1("direct")

	go ft1("goroutine") //协程

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

// Result
// direct : 0
// direct : 1
// direct : 2
// goroutine : 0
// going
// goroutine : 1
// goroutine : 2
// <enter>
// done
