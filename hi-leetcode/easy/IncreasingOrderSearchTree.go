package easy

import "fmt"

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
