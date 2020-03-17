package easy

import (
	"container/list"
	"testing"
)

func buildTree() *Node {
	/**
	 *       5
	 *      / \
	 *     4   8
	 *    /   / \
	 *   11  13  4
	 *  /  \    / \
	 * 7    2  5   1
	 *
	 */
	_7 := NewNode(7, nil, nil)
	_2 := NewNode(2, nil, nil)
	_1 := NewNode(1, nil, nil)
	_5 := NewNode(5, nil, nil)

	_11 := NewNode(11, _7, _2)
	_13 := NewNode(13, nil, nil)
	_4 := NewNode(4, _5, _1)

	_4u := NewNode(4, _11, nil)
	_8 := NewNode(8, _13, _4)
	_5u := NewNode(5, _4u, _8)
	return _5u
}

func TestHasPathSum(t *testing.T) {
	res := HasPathSum(buildTree(), 22)
	println(res)
}

func TestHasPathSum1(t *testing.T) {
	res := HasPathSum1(buildTree(), 22)
	println(res)
}

func TestGetPathSum(t *testing.T) {
	res := GetPathSum(buildTree(), 22)
	tmp := res.Front()
	for nil != tmp {
		tmpL := tmp.Value.(*list.List)
		cur := tmpL.Front()
		for nil != cur {
			print(cur.Value.(int))
			print(" ")
			cur = cur.Next()
		}
		print("\n")
		tmp = tmp.Next()
	}
}

func TestGetPathSum1(t *testing.T) {
	res := GetPathSum1(buildTree(), 22)
	tmp := res.Front()
	for nil != tmp {
		tmpL := tmp.Value.(*list.List)
		cur := tmpL.Front()
		for nil != cur {
			print(cur.Value.(int))
			print(" ")
			cur = cur.Next()
		}
		print("\n")
		tmp = tmp.Next()
	}
}

func TestOtherPathSum(t *testing.T) {
	arr := []int{113, 215, 221}
	result := OtherPathSum(arr)
	println(result)
}
