package structure

import (
	"container/list"
	"testing"
)

func levelPrint(node *TreeNode) {
	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		for size, i := queue.Len(), 0; i < size; i++ {
			front := queue.Front()
			cur := front.Value.(*TreeNode)
			queue.Remove(front)
			print(cur.Value.(int))
			print(" ")
			if nil != cur.Left {
				queue.PushBack(cur.Left)
			}
			if nil != cur.Right {
				queue.PushBack(cur.Right)
			}
		}
		println()
	}
}
func TestReConstructBinaryTree(t *testing.T) {
	pre, in := []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}
	node := ReConstructBinaryTree(pre, in)
	levelPrint(node)
}

func prepareTree1() *TreeNode {
	/**
	 *      1
	 *    /  \
	 *   2    3
	 *  / \
	 * 4   5
	 */

	_5 := &TreeNode{Value: 5, Left: nil, Right: nil}
	_4 := &TreeNode{Value: 4, Left: nil, Right: nil}
	_3 := &TreeNode{Value: 3, Left: nil, Right: nil}
	_2 := &TreeNode{Value: 2, Left: _4, Right: _5}
	_1 := &TreeNode{Value: 1, Left: _2, Right: _3}
	return _1
}
func prepareTree2() *TreeNode {
	/**
	 *   2
	 *  / \
	 * 4   5
	 */
	_5 := &TreeNode{Value: 5, Left: nil, Right: nil}
	_4 := &TreeNode{Value: 4, Left: nil, Right: nil}
	_2 := &TreeNode{Value: 2, Left: _4, Right: _5}
	return _2
}
func prepareBST1() *TreeNode {
	/**
	        4
	      /   \
	    2      6
	   / \    / \
	  1   3  5   7
	*/

	_1 := &TreeNode{Value: 1, Left: nil, Right: nil}
	_3 := &TreeNode{Value: 3, Left: nil, Right: nil}
	_5 := &TreeNode{Value: 5, Left: nil, Right: nil}
	_7 := &TreeNode{Value: 7, Left: nil, Right: nil}

	_2 := &TreeNode{Value: 2, Left: _1, Right: _3}
	_6 := &TreeNode{Value: 6, Left: _5, Right: _7}

	_4 := &TreeNode{Value: 4, Left: _2, Right: _6}
	return _4
}
func TestIsSameTree(t *testing.T) {
	n1, n2 := prepareTree2(), prepareTree2()
	res := IsSameTree(n1, n2)
	println(res)
}
func TestIsSubtree(t *testing.T) {
	n1, n2 := prepareTree1(), prepareTree2()
	res := IsSubtree(n1, n2)
	println(res)
}
func TestIsSubtree0(t *testing.T) {
	n1, n2 := prepareTree1(), prepareTree2()
	res := IsSubtree0(n1, n2)
	println(res)
}

func TestMirrorTree0(t *testing.T) {
	n1 := prepareTree1()
	MirrorTree0(n1)
	levelPrint(n1)
}
func TestMirrorTree1(t *testing.T) {
	n1 := prepareTree1()
	MirrorTree1(n1)
	levelPrint(n1)
}

func TestPrintFromTopToBottom(t *testing.T) {
	n1 := prepareTree1()
	PrintFromTopToBottom(n1)
}

func TestVerifySequenceOfBST(t *testing.T) {
	postOrder := []int{1, 2, 3, 4, 5}
	res := VerifySequenceOfBST(postOrder)
	println(res)
}

func TestFindPathInBinaryTree(t *testing.T) {
	n1 := prepareBST1()
	res := FindPathInBinaryTree(n1, 9)
	if nil == res {
		return
	}
	cur := res.Front()
	for nil != cur {
		inner := cur.Value.(*list.List)
		innerCur := inner.Front()
		for nil != innerCur {
			print(innerCur.Value.(int))
			print(" ")

			innerCur = innerCur.Next()
		}

		cur = cur.Next()
	}
	println()
}

func TestConvertTree2LinkedList(t *testing.T) {
	n1 := prepareBST1()
	res := ConvertTree2LinkedList(n1)
	for nil != res {
		print(res.Value.(int))
		print(" ")
		res = res.Right
	}
	println()
}

func TestDepthOfBinaryTree(t *testing.T) {
	n1 := prepareBST1()
	res := DepthOfBinaryTree(n1)
	println(res)
}

func TestIsBalanceBinaryTree(t *testing.T) {
	n1 := prepareBST1()
	res := IsBalanceBinaryTree(n1)
	println(res)
}

func TestGetInOrderNextInBinaryTree(t *testing.T) {
	/**
	 *      1
	 *    /  \
	 *   2    3
	 *  / \
	 * 4   5
	 */
	_5 := &TreeNode{Value: 5, Left: nil, Right: nil}
	_4 := &TreeNode{Value: 4, Left: nil, Right: nil}
	_3 := &TreeNode{Value: 3, Left: nil, Right: nil}
	_2 := &TreeNode{Value: 2, Left: _4, Right: _5}
	_1 := &TreeNode{Value: 1, Left: _2, Right: _3}
	println(_1)

	res := GetInOrderNextInBinaryTree(_2)
	if nil != res {
		println(res.Value.(int))
	}
}

func TestIsSymmetricalTree(t *testing.T) {
	/**
	 *         1
	 *       /  \
	 *      2    2
	 *     / \  / \
	 *    6  7 7   6
	 */
	_7u := &TreeNode{Value: 7, Left: nil, Right: nil}
	_6u := &TreeNode{Value: 6, Left: nil, Right: nil}
	_7 := &TreeNode{Value: 7, Left: nil, Right: nil}
	_6 := &TreeNode{Value: 6, Left: nil, Right: nil}
	_2u := &TreeNode{Value: 2, Left: _7u, Right: _6u}
	_2 := &TreeNode{Value: 2, Left: _6, Right: _7}
	_1 := &TreeNode{Value: 1, Left: _2, Right: _2u}

	res := IsSymmetricalTree(_1)
	println(res)
}

func TestZigzagLevelOrderPrint(t *testing.T) {
	n1 := prepareTree1()
	ZigzagLevelOrderPrint(n1)
}

func TestLevelOrderPrint(t *testing.T) {
	n1 := prepareTree1()
	LevelOrderPrint(n1)
}

func TestKthInBST(t *testing.T) {
	n1 := prepareBST1()
	res := KthInBST(n1, 3)
	if nil != res {
		println(res.Value.(int))
	}
}
