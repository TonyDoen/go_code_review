package medium

import (
	"fmt"
	"testing"
)

func TestDeleteNode3(t *testing.T) {
	/**
	        7
	      /   \
	     4     8
	   /   \
	  2     6
	   \    /
	    3  5
	*/
	nd5 := NewNode(5, nil, nil)
	nd3 := NewNode(3, nil, nil)
	nd2 := NewNode(2, nil, nd3)
	nd6 := NewNode(6, nd5, nil)
	nd4 := NewNode(4, nd2, nd6)
	nd8 := NewNode(8, nil, nil)
	nd7 := NewNode(7, nd4, nd8)

	res := DeleteNode3(nd7, 4)

	fmt.Printf("\ni am preOrder\n")
	res.preOrder()
	fmt.Printf("\ni am inOrder\n")
	res.inOrder()
	fmt.Printf("\ni am postOrder\n")
	res.postOrder()
	fmt.Printf("\ni am levelOrder\n")
	res.levelOrder()
}
