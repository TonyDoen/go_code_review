package main

import "fmt"

func at() {
	fmt.Println("func at()")
}

func bt() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in bt()")
		}
	}()
	panic("panic in bt()")
	fmt.Println("func bt()")
}

func ct() {
	fmt.Println("func ct()")
}

func main() {
	at()
	bt()
	ct()
}
