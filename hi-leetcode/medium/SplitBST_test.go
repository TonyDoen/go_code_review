package medium

import (
	"fmt"
	"testing"
)

func TestSplitBST1 (t *testing.T) {
	/**
	        4
	      /   \
	    2      6
	   / \    / \
	  1   3  5   7
	*/
	nd1 := NewNode(1, nil, nil)
	nd3 := NewNode(3, nil, nil)
	nd5 := NewNode(5, nil, nil)
	nd7 := NewNode(7, nil, nil)
	nd2 := NewNode(2, nd1, nd3)
	nd6 := NewNode(6, nd5, nd7)
	nd4 := NewNode(4, nd2, nd6)

	res := SplitBST1(nd4, 2)
	for i := 0; i < 2; i++ {
		fmt.Printf("\ni am %d preOrder\n", i)
		res[i].preOrder()
		fmt.Printf("\ni am %d inOrder\n", i)
		res[i].inOrder()
		fmt.Printf("\ni am %d postOrder\n", i)
		res[i].postOrder()
	}
}