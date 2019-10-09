package medium

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	_3 := &lNode{data: 3}
	_4 := &lNode{data: 4, next: _3}
	_2 := &lNode{data: 2, next: _4}

	_4_1 := &lNode{data: 4}
	_6 := &lNode{data: 6, next: _4_1}
	_5 := &lNode{data: 5, next: _6}

	res := AddTwoNumbers1(_2, _5)
	for tmp := res; nil != tmp; {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
}
