package easy

import (
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	root := prepareTree()
	res := IncreasingBST(root)
	res.preOrder()
}

func prepareTree() *Node {
	/**
	      5
	     / \
	    3   6
	   / \   \
	  2   4   8
	 /       / \
	1       7   9
	*/
	nd1 := NewNode(1, nil, nil)
	nd7 := NewNode(7, nil, nil)
	nd9 := NewNode(9, nil, nil)
	nd2 := NewNode(2, nd1, nil)
	nd4 := NewNode(4, nil, nil)
	nd8 := NewNode(8, nd7, nd9)
	nd3 := NewNode(3, nd2, nd4)
	nd6 := NewNode(6, nil, nd8)
	nd5 := NewNode(5, nd3, nd6)
	return nd5
}
