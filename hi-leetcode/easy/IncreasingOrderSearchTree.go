package easy

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data        int
	Left, Right *Node
}
// 树的 前序遍历; 中序遍历; 后序遍历
func (n *Node) preOrder() { // 前序遍历
	if nil == n {
		return
	}
	fmt.Printf("%d ", n.Data) // visit
	n.Left.preOrder()
	n.Right.preOrder()
}
func NewNode(d int, l, r *Node) *Node {
	return &Node{Data: d, Left: l, Right: r}
}

func IncreasingBST(root *Node) *Node {
	return iBST(root, nil)
}
func iBST(node, pre *Node) *Node {
	if nil == node {
		return pre
	}
	res := iBST(node.Left, node)
	node.Left = nil
	node.Right = iBST(node.Right, pre)
	return res
}

func IncreasingBST2(root *Node) *Node {
	dummy := NewNode(-1, nil, nil)
	pre := dummy
	var st list.List
	for st.Len() > 0 || nil != root {
		for nil != root {
			st.PushFront(root)
			root = root.Left
		}

		el := st.Front()
		root = el.Value.(*Node)
		st.Remove(el)

		pre.Right = root
		pre = pre.Right
		root.Left = nil
		root = root.Right
	}
	return dummy.Right
}
