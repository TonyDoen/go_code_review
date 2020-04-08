package medium

import (
	"container/list"
	"testing"
)

func TestDifferentWaysToCompute(t *testing.T) {
	input := "2*3-4*5"
	res := DifferentWaysToCompute(input)
	printList(res)
}

func TestDifferentWaysToCompute1(t *testing.T) {
	input := "2*3-4*5"
	res := DifferentWaysToCompute1(input)
	printList(res)
}

func TestDifferentWaysToCompute2(t *testing.T) {
	input := "2*3-4*5"
	res := DifferentWaysToCompute2(input)
	printList(res)
}

func printList(res *list.List) {
	print("[")
	for cur := res.Front(); nil != cur; {
		print(cur.Value.(int))
		cur = cur.Next()
		if nil != cur {
			print(", ")
		}
	}
	println("]")
}
