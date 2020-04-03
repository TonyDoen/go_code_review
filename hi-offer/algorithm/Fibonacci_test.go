package algorithm

import "testing"

func TestFibonacci1(t *testing.T) {
	idx := Fibonacci1(3)
	println(idx)
}

func TestFibonacci2(t *testing.T) {
	idx := Fibonacci2(0)
	println(idx)
}

func TestFibonacci3(t *testing.T) {
	res := Fibonacci3(6)
	println(res)
}

func TestJump(t *testing.T) {
	n := 5
	res := Jump(n)
	println(res)
}

func TestFreakJump1(t *testing.T) {
	n := 5
	res := FreakJump1(n)
	println(res)

	res = FreakJump2(n)
	println(res)
}

func TestRectCover(t *testing.T) {
	n := 3
	res := RectCover(n)
	println(res)
}
