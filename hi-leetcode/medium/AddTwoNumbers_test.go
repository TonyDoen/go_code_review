package medium

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	_3 := &lNode{data: 3}
	_4 := &lNode{data: 4, next: _3}
	_2 := &lNode{data: 2, next: _4}

	_4u := &lNode{data: 4}
	_6 := &lNode{data: 6, next: _4u}
	_5 := &lNode{data: 5, next: _6}

	res := AddTwoNumbers1(_2, _5)
	for tmp := res; nil != tmp; {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
}

func TestAddTwoNumber2(t *testing.T) {
	// (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	_3 := &lNode{data: 3}
	_4 := &lNode{data: 4, next: _3}
	_2 := &lNode{data: 2, next: _4}
	_7 := &lNode{data: 7, next: _2}

	_4u := &lNode{data: 4}
	_6 := &lNode{data: 6, next: _4u}
	_5 := &lNode{data: 5, next: _6}
	res := AddTwoNumber2(_7, _5)
	for nil != res {
		print(res.data)
		res = res.next
	}
	println()
}
