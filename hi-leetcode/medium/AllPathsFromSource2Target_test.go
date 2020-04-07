package medium

import (
	"container/list"
	"testing"
)

func TestAllPathsFromSource2Target0(t *testing.T) {
	/**
	 * [[1,2], [3], [3], []]
	 *
	 * 0--->1
	 * |    |
	 * v    v
	 * 2--->3
	 */
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	res := AllPathsFromSource2Target0(graph)
	print2List(res)
}

func TestAllPathsFromSource2Target1(t *testing.T) {
	/**
	 * [[1,2], [3], [3], []]
	 *
	 * 0--->1
	 * |    |
	 * v    v
	 * 2--->3
	 */
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	res := AllPathsFromSource2Target1(graph)
	print2List(res)
}

func print2List(res *list.List) {
	front := res.Front()
	for nil != front {
		lt := front.Value.(*list.List)
		ft := lt.Front()
		for nil != ft {
			print(ft.Value.(int))
			ft = ft.Next()
			if nil != ft {
				print("->")
			}
		}
		front = front.Next()
		println()
	}
}
