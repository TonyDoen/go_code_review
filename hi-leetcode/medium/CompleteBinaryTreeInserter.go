package medium

import (
	"container/list"
	"strconv"
)

/**
 * Complete Binary Tree Inserter
 *
 * A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.
 * Write a data structure CBTInserter that is initialized with a complete binary tree and supports the following operations:
 * 1. CBTInserter(TreeNode root) initializes the data structure on a given tree with head node root;
 * 2. CBTInserter.insert(int v) will insert a TreeNode into the tree with value node.val = v so that the tree remains complete, and returns the value of the parent of the inserted TreeNode;
 * 3. CBTInserter.get_root() will return the head node of the tree.
 *
 * Example 1:
 * Input: inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]]
 * Output: [null,1,[1,2]]
 *
 * Example 2:
 * Input: inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
 * Output: [null,3,4,[1,2,3,4,5,6,7,8]]
 *
 * Note:
 * 1. The initial given tree is complete and contains between 1 and 1000 nodes.
 * 2. CBTInserter.insert is called at most 10000 times per test case.
 * 3. Every value of a given or inserted node is between 0 and 5000.
 *
 *
 * 题意：完全二叉树插入器
 */
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (t *TreeNode) Print() {
	lt := list.New()
	lt.PushBack(t)
	for lt.Len() > 0 {
		front := lt.Front()
		node := front.Value.(*TreeNode)
		print("{Value:" + strconv.Itoa(node.Value))
		if nil != node.Left {
			lt.PushBack(node.Left)
			print("; Left:" + strconv.Itoa(node.Left.Value))
		}
		if nil != node.Right {
			lt.PushBack(node.Right)
			print("; Right:" + strconv.Itoa(node.Right.Value))
		}
		println("}")
		lt.Remove(front)
	}
}

type CompleteBinaryTreeInserter struct {
	Root  *TreeNode
	Deque *list.List
}

func NewCompleteBinaryTreeInserter(root *TreeNode) *CompleteBinaryTreeInserter {
	result, queue := &CompleteBinaryTreeInserter{Root: root, Deque: list.New()}, list.New()

	for queue.PushBack(root); queue.Len() > 0; {
		front := queue.Front()
		node := front.Value.(*TreeNode)
		queue.Remove(front)

		if nil == node.Left || nil == node.Right {
			result.Deque.PushBack(node)
		}
		if nil != node.Left {
			result.Deque.PushBack(node.Left)
		}
		if nil != node.Right {
			result.Deque.PushBack(node.Right)
		}
	}
	return result
}

func (c *CompleteBinaryTreeInserter) Insert(v int) int {
	front := c.Deque.Front()
	node := front.Value.(*TreeNode)
	result := node.Value
	newNode := &TreeNode{Value: v}
	c.Deque.PushBack(newNode)
	if nil == node.Left {
		node.Left = newNode
	} else {
		node.Right = newNode
		c.Deque.Remove(front) // remove first
	}
	return result
}

func (c *CompleteBinaryTreeInserter) GetRoot() *TreeNode {
	return c.Root
}

//type CBTInserter struct {
//	root             *TreeNode
//	LastLastNodeList []*TreeNode
//	LastNodeList     []*TreeNode
//}
//
//func Constructor(root *TreeNode) CBTInserter {
//	this := CBTInserter{}
//	this.root = root
//
//	nodeList := []*TreeNode{root}
//	for len(nodeList) > 0 {
//		l := len(nodeList)
//		this.LastLastNodeList = this.LastNodeList
//		this.LastNodeList = nodeList
//		for i := 0; i < l; i++ {
//			nodeTmp := nodeList[0]
//			nodeList = nodeList[1:]
//
//			if nodeTmp.Left != nil {
//				nodeList = append(nodeList, nodeTmp.Left)
//			}
//			if nodeTmp.Right != nil {
//				nodeList = append(nodeList, nodeTmp.Right)
//			}
//		}
//	}
//	return this
//}
//
//func (this *CBTInserter) Insert(v int) int {
//	if len(this.LastNodeList) >= len(this.LastLastNodeList)*2 {
//		this.LastLastNodeList = this.LastNodeList
//		this.LastNodeList = []*TreeNode{}
//	}
//	newNode := &TreeNode{Value: v}
//	l := len(this.LastNodeList)
//	if l%2 == 0 {
//		this.LastLastNodeList[l/2].Left = newNode
//	} else {
//		this.LastLastNodeList[l/2].Right = newNode
//	}
//	this.LastNodeList = append(this.LastNodeList, newNode)
//	return this.LastLastNodeList[l/2].Value
//}
//
//func (this *CBTInserter) GetRoot() *TreeNode {
//	return this.root
//}
