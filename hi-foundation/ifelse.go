package main

import "fmt"

func main() {
	if 0 == 7%2 {
		fmt.Println("0 == 7%2")
	} else {
		fmt.Println("0 != 7%2")
	}

	if 0 == 7%2 {
		fmt.Println("0 == 7%2")
	} else if 0 == 8%2 {
		fmt.Println("0 == 8%2")
	} else {
		fmt.Println("0 != 7%2")
	}
}
