package main

import (
	"fmt"
	"math"
)

// const
const (
	str1 string = "constant str1"
	str2 string = "constant str2"
)

// type alias
type INT int

func (i *INT) PrintStarINT() {
	fmt.Println("INT", &i, i, *i)
}
func (i INT) PrintINT() {
	fmt.Println("INT", &i, i)
}

var testINT INT = 99

// HA := "High Avaible" // error
var HA string = "High Avaible" // bingo

func main() {
	fmt.Println(str1, str2)
	// str1 = "hahaha" // error
	const (
		n = 5000000
		d = 3e20 / n
	)

	const a1 = 66
	var a2 int = 66
	a3 := 66

	fmt.Println(a1, a2, a3)
	fmt.Println(n, d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	testINT.PrintStarINT()
	testINT.PrintINT()
}
