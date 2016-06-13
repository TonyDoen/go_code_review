package main

import "fmt"

// Command
// # go run helloworld.go

func main() {
	// value
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// variables
	var a string = "this is string."
	var b, c int = 1, 2
	var d = true
	var e int

	f := "short"
	fmt.Println(a, b, c, d, e, f)

	//pointer
	var x int = 4
	y := 4

	var v *int = new(int) //return pointer, clear mem 0, not init mem
	*v = 4
	w := new(int)
	*w = 4
	fmt.Println(x, y, v, w)
	fmt.Println(&x, &y, *v, *w)

	// new和make的区别
	// new 是一个分配内存的内建函数, golang 它只是将内存清零，而不是初始化内存. new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值. 返回了一个指向新分配的类型为T的零值的指针.
	// make(T, args)函数的目的与new(T)不同. 它仅用于创建切片、map和chan（消息管道）. 返回值类型T（不是*T）. 返回值被初始化了的（不是零）实例. args包括指向数据（在一个数组内部）的指针、长度以及容量,在这三项内容被初始化之前，切片值为nil.
	var p *[]int = new([]int)     // 为切片结构分配内存；*p == nil；很少使用
	var q []int = make([]int, 10) // 切片q现在是对一个新的有10个整数的数组的引用
	fmt.Println(p, q)

	*p = make([]int, 10, 10)
	fmt.Println(p, *p, (*p)[2])

	var p2 []int = make([]int, 10)
	usualArr := make([]int, 10)
	fmt.Println(p2, usualArr)
}
