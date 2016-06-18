package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) walkWith(pwith *Person) {
	p.name = "walkWith"
	fmt.Println("hello, I am", p.name, ". I am walking with", pwith.name)
}
func (p *Person) walkAlongWith(pwith *Person) {
	p.name = "walkAlongWith"
	fmt.Println("hello, I am", p.name, ". I am walking along with", pwith.name)

	a := &struct { // 匿名结构
		name, flag string
		age        int
		contact    struct{ originPhone, originCity string } // 嵌套结构
	}{
		name: "robber",
		age:  55,
		flag: "bad-person",
	}
	a.contact.originCity = "china-city" // 嵌套结构 init
	a.contact.originPhone = "110"
	fmt.Println("now bad thing has happened. we have been robbed by", a)

	b := &struct {
		string // 匿名字段
		// string // error: duplicate field string
		int
	}{
		"police",
		911,
	}
	fmt.Println(b, " has got the robber.")

	confictc := &struct {
		Person
		name string // confict
	}{
		Person: Person{"TonyStack", 34},
		name:   "NewName",
	}
	fmt.Println(confictc, confictc.name, confictc.Person.name)
}

func main() {
	fmt.Println(Person{"bob", 20})
	fmt.Println(Person{name: "bob", age: 20})
	fmt.Println(Person{name: "bob"})
	fmt.Println(Person{age: 20})

	p1 := Person{name: "tony", age: 32}
	p1.age = 22
	fmt.Println(&p1, p1, p1.name, p1.age)

	p1.walkWith(&Person{name: "doen", age: 55})
	fmt.Println(&p1, p1, p1.name, p1.age) // p1.name no change
	// vs.
	p1.walkAlongWith(&Person{name: "doen", age: 55})
	fmt.Println(&p1, p1, p1.name, p1.age) // p1.name changed
}
