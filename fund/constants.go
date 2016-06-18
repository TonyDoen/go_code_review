package main

import (
	"fmt"
	"math"
)

const (
	str1 string = "constant str1"
	str2 string = "constant str2"
)

func main() {
	fmt.Println(str1, str2)
	// str1 = "hahaha" // error
	const (
		n = 5000000
		d = 3e20 / n
	)
	fmt.Println(n, d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
