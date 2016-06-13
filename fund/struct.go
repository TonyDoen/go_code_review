package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"bob", 20})
	fmt.Println(Person{name: "bob", age: 20})
	fmt.Println(Person{name: "bob"})
	fmt.Println(Person{age: 20})

	p1 := Person{name: "tony", age: 32}
	p1.age = 22
	fmt.Println(p1, p1.name, p1.age)
}
