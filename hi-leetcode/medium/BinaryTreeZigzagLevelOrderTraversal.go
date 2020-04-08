package medium

import "container/list"

/**
 * 059-按之字形顺序打印二叉树
 *
 * 题意：按照z字形层次遍历二叉树（以根节点所在层为第1层，则第二层的变量从右边节点开始直到最左边节点，第三层遍历则是从最左边开始到最右边）
 * 思路：z字形层次遍历是对层次遍历加上了一个限制条件（即相邻层，从左到右的遍历顺序相反），因此还是可以采用队列来实现，只不过节点接入队列时需要考虑加入的顺序
 */
func ZigzagLevelOrderPrint(root *TreeNode) {
	if nil == root {
		return
	}

	queue := list.New()
	queue.PushBack(root)
	for depth := 0; queue.Len() > 0; depth++ {
		for i, size := 0, queue.Len(); i < size; i++ {
			if 0 == depth%2 {
				front := queue.Front()
				node := front.Value.(*TreeNode)
				print(node.Value)
				print(" ")
				queue.Remove(front)

				if nil != node.Left {
					queue.PushBack(node.Left)
				}
				if nil != node.Right {
					queue.PushBack(node.Right)
				}
			} else {
				back := queue.Back()
				node := back.Value.(*TreeNode)
				print(node.Value)
				print(" ")
				queue.Remove(back)

				if nil != node.Right {
					queue.PushFront(node.Right)
				}
				if nil != node.Left {
					queue.PushFront(node.Left)
				}
			}
		}
		println()
	}
}
