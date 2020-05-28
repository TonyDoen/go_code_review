package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Println(m1)

	m1["key1"] = 1234
	m1["key2"] = 5678
	fmt.Println("map:", m1, "len(m1)", len(m1))

	v1 := m1["key1"]
	fmt.Println(v1)

	delete(m1, "key1")
	fmt.Println(m1)

	_, prs := m1["key2"]
	fmt.Println(prs)
	_, prs = m1["key3"]
	fmt.Println(prs)

	m2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(m2)
}
