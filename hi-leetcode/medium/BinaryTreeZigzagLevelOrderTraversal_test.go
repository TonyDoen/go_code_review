package medium

import "testing"

func TestZigzagLevelOrderPrint(t *testing.T) {
	/**
	        7
	      /   \
	     4     8
	   /   \
	  2      6
	   \    / \
	    3  5   9
	*/
	_9 := &TreeNode{Value: 9, Left: nil, Right: nil}
	_5 := &TreeNode{Value: 5, Left: nil, Right: nil}
	_3 := &TreeNode{Value: 3, Left: nil, Right: nil}
	_6 := &TreeNode{Value: 6, Left: _5, Right: _9}
	_2 := &TreeNode{Value: 2, Left: nil, Right: _3}
	_8 := &TreeNode{Value: 8, Left: nil, Right: nil}
	_4 := &TreeNode{Value: 4, Left: _2, Right: _6}
	_7 := &TreeNode{Value: 7, Left: _4, Right: _8}

	ZigzagLevelOrderPrint(_7)
}
