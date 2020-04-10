package medium

import (
	"testing"
)

func TestSortLinkedList0(t *testing.T) {
	// 7 -> 2 -> 4 -> 3
	_3 := &lNode{data: 3}
	_4 := &lNode{data: 4, next: _3}
	_2 := &lNode{data: 2, next: _4}
	_7 := &lNode{data: 7, next: _2}
	res := SortLinkedList0(_7)
	for nil != res {
		print(res.data)
		res = res.next
		if nil != res {
			print("->")
		}
	}
	println()
}

func TestSortLinkedList1(t *testing.T) {
	// 7 -> 2 -> 4 -> 3
	_3 := &lNode{data: 3}
	_4 := &lNode{data: 4, next: _3}
	_2 := &lNode{data: 2, next: _4}
	_7 := &lNode{data: 7, next: _2}
	res := SortLinkedList1(_7)
	for nil != res {
		print(res.data)
		res = res.next
		if nil != res {
			print("->")
		}
	}
	println()
}
