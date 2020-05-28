package main

import "fmt"

func main() {
	arr1 := []int{11, 12, 31, 45}
	sum := 0
	for _, num := range arr1 {
		sum += num
	}
	fmt.Println(sum)

	for index, value := range arr1 {
		fmt.Println(index, value)
	}

	kv1 := map[string]string{"key1": "apple", "key2": "banana"}
	for k, v := range kv1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "golang" {
		fmt.Printf("%d -> %c, %d\n", i, c, c)
	}
}
