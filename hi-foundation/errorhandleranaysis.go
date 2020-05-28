package main

import (
	"fmt"
)

func returntripleint() (a, b, c int) {
	a, b, c = 1, 2, 3
	return a, b, c
}

func printunkownargs(b string, a ...int) {
	fmt.Println("b = ", b)
	fmt.Println("a = ", a)
}

func main() {
	fmt.Println(returntripleint())
	printunkownargs("this is b", 2, 5, 1, 3, 1, 4)

	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}
	for _, f := range fs {
		f()
	}
}
