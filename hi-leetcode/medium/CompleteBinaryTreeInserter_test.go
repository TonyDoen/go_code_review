package medium

import "testing"

func TestCompleteBinaryTreeInserter(t *testing.T) {
	root := &TreeNode{Value: 1}
	inserter := NewCompleteBinaryTreeInserter(root)
	for i := 10; i > 1; i-- {
		inserter.Insert(i)
	}
	inserter.GetRoot().Print()
}
