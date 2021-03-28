package main

import "fmt"

func sum1(x, y int) int {
	return x + y
}

func sum2(x, y, z int) int {
	return x + y + z
}

func initvals() (int, int, int) { // multiple return values.
	return 1, 2, 3
}

func sum3(x ...int) int { // variadic functions
	// fmt.Println(x)
	res := 0
	for _, num := range x {
		res += num
	}
	return res
}

func intseq() func() int { // anonymous functions,
	i := 0
	return func() int { // Anonymous functions are useful when you want to define a function inline without having to name it.
		i += 1
		return i
	}
}

func recurfact(n int) int { // recursion
	if 0 == n {
		return 1
	}
	return n * recurfact(n-1)
}

func usepointer1(ival int) int { // use pointer
	ival = 0
	return ival
}
func usepointer2(iptr *int) *int {
	*iptr = 0
	return iptr
}
func usepointer3(iptr *int) int {
	*iptr = 0
	return *iptr
}
func usepointer4(iptr *int) {
	*iptr = 0
}

func main() {
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(3, 4, 5))
	fmt.Println(initvals())
	fmt.Println(sum3(3, 4, 5, 6, 7, 8))
	fmt.Println(sum3([]int{1, 23}...)) // haha...

	testf := intseq()
	fmt.Println(intseq, intseq(), testf, testf())

	fmt.Println(recurfact(4))

	testv := 44
	fmt.Println("pointer:\n", usepointer1(testv), testv)
	fmt.Println(usepointer2(&testv), testv)
	fmt.Println(usepointer3(&testv), testv)
	usepointer4(&testv)
	fmt.Println(testv)
}
