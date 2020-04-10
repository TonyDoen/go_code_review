package medium

import (
	"container/list"
	"testing"
)

func TestGetSubsets00(t *testing.T) {
	arr := []int{1, 2, 3}
	res := GetSubsets00(arr)
	print2IntArray(res)
}

func TestGetSubsets10(t *testing.T) {
	arr := []int{1, 2, 2}
	res := GetSubsets10(arr)
	print2IntArray(res)
}

func print2IntArray(res *list.List) {
	print("[")
	front := res.Front()
	for nil != front {
		lt := front.Value.(*list.List)
		ft := lt.Front()
		print("[")
		for nil != ft {
			print(ft.Value.(int))
			ft = ft.Next()
			if nil != ft {
				print(",")
			}
		}
		front = front.Next()
		print("]")

		if nil != front {
			print(",")
		}
	}
	println("]")
}
