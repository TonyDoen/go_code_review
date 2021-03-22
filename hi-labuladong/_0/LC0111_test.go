package _0

import (
	"fmt"
	"testing"
)

func TestMinDepth(t *testing.T) {
	/**
	 *         3
	 *       /   \
	 *     9     20
	 *           /  \
	 *          15   7
	 */
	_15 := &TreeNode{15, nil, nil}
	_7 := &TreeNode{7, nil, nil}
	_20 := &TreeNode{20, _15, _7}
	_9 := &TreeNode{9, nil, nil}
	_3 := &TreeNode{3, _9, _20}

	rs := MinDepth(_3)
	fmt.Printf("%d\n", rs)
}
