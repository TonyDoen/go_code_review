package structure

import (
	"container/list"
	"testing"
)

func prepareLinkedList() *Node {
	/**
	 * 4 -> 3 -> 2 -> 1
	 */
	_1 := NewNode(1, nil)
	_2 := NewNode(2, _1)
	_3 := NewNode(3, _2)
	_4 := NewNode(4, _3)
	return _4
}

func TestPrintSingleLinkedList(t *testing.T) {
	head := prepareLinkedList()
	PrintSingleLinkedList(head)
}

func TestFindKth2Tail(t *testing.T)  {
	head := prepareLinkedList()
	res := FindKth2Tail(head, 2)
	println(res.Value.(int))
}

func TestReverseSingleLinkedList(t *testing.T)  {
	head := prepareLinkedList()
	res := ReverseSingleLinkedList(head)
	for nil != res {
		print(res.Value.(int))
		print(" ")
		res = res.Next
	}
	println()
}

func TestMergeTwoNode(t *testing.T)  {
	n1 := prepareLinkedList()
	n2 := prepareLinkedList()
	res := MergeTwoNode(n1, n2)
	for nil != res {
		print(res.Value.(int))
		print(" ")
		res = res.Next
	}
	println()
}

func TestMergeKthNode(t *testing.T)  {
	lt := list.New()
	for i := 0; i < 3; i++ {
		lt.PushBack(prepareLinkedList())
	}
	res := MergeKthNode(lt)
	for nil != res {
		print(res.Value.(int))
		print(" ")
		res = res.Next
	}
	println()
}

func TestComplexNode_Clone(t *testing.T) {
	/**
	 * 4 -> 3 -> 2 -> 1
	 */
	_1 := &ComplexNode{Value: 1, Next: nil}
	_2 := &ComplexNode{Value: 2, Next: _1}
	_3 := &ComplexNode{Value: 3, Next: _2}
	_4 := &ComplexNode{Value: 4, Next: _3}
	_1.Random = _3
	_2.Random = _4
	_3.Random = _2
	_4.Random = _1

	cur := _4.Clone()
	for nil != cur {
		print("{curValue:")
		print(cur.Value.(int))
		print("}")
		cur = cur.Next
	}
	println()
}

func TestFindFirstCommonNode(t *testing.T) {
	/**
	 * Y型:
	 * 1 -> 2 -> 3
	 *            \
	 *             6 -> 7 -> 8
	 *            /
	 *      4 -> 5
	 *
	 */
	_8 := NewNode(8, nil)
	_7 := NewNode(7, _8)
	_6 := NewNode(6, _7)
	_5 := NewNode(5, _6)
	_4 := NewNode(4, _5)

	_3 := NewNode(3, _6)
	_2 := NewNode(2, _3)
	_1 := NewNode(1, _2)

	res := FindFirstCommonNode(_1, _4)
	if nil != res {
		println(res.Value.(int))
	}
}

func TestFindEntryNodeOfLoop(t *testing.T) {
	/**
	 * 1 -> 2 -> 3 -> 4 -> 5 -> 6
	 *                |         |
	 *                9 <- 8 <- 7
	 */
	_9 := NewNode(9, nil)
	_8 := NewNode(8, _9)
	_7 := NewNode(7, _8)
	_6 := NewNode(6, _7)
	_5 := NewNode(5, _6)
	_4 := NewNode(4, _5)
	_3 := NewNode(3, _4)
	_2 := NewNode(2, _3)
	_1 := NewNode(1, _2)
	// loop
	_9.Next = _4

	res := FindEntryNodeOfLoop(_1)
	if nil != res {
		println(res.Value.(int))
	}
}

func prepareDuplicateLinkedList() *Node {
	/**
	 * 5 -> 4 -> 4 -> 3 -> 3 -> 2 -> 1 -> 1
	 */
	_1 := NewNode(1, nil)
	_1u := NewNode(1, _1)
	_2 := NewNode(2, _1u)
	_3 := NewNode(3, _2)
	_3u := NewNode(3, _3)
	_4 := NewNode(4, _3u)
	_4u := NewNode(4, _4)
	_5 := NewNode(5, _4u)
	return _5
}
func TestDeleteDuplicateNode(t *testing.T) {
	_5 := prepareDuplicateLinkedList()

	res := DeleteDuplicateNode(_5)
	for nil != res {
		print(res.Value.(int))
		print(" ")
		res = res.Next
	}
	println()
}

func TestDeleteDuplicateNodeRemainFirstOne(t *testing.T) {
	_5 := prepareDuplicateLinkedList()

	res := DeleteDuplicateNodeRemainFirstOne(_5)
	for nil != res {
		print(res.Value.(int))
		print(" ")
		res = res.Next
	}
	println()
}