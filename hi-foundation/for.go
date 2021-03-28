package main

import "fmt"

func main() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i += 1
		i++ //not available for python
	}

	for i = 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	for {
		fmt.Print("\n" + "ever loop")
		break
	}
}
