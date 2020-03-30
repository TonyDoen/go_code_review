package structure

import "container/list"

type TreeNode struct {
	Value               interface{}
	Left, Right, Parent *TreeNode
}

/**
 * 题目描述
 * 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
 * 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
 *
 * 先来分析一下前序遍历和中序遍历得到的结果，
 * 前序遍历第一位是根节点；
 * 中序遍历中，根节点左边的是根节点的左子树，右边是根节点的右子树。
 * 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}。
 *
 * 首先，根节点 是{ 1 }；
 * 左子树是：前序{ 2,4,7 } ，中序{ 4,7,2 }；
 * 右子树是：前序{ 3,5,6,8 } ，中序{ 5,3,8,6 }；
 * 这时，如果我们把左子树和右子树分别作为新的二叉树，则可以求出其根节点，左子树和右子树。
 * 这样，一直用这个方式，就可以实现重建二叉树。
 */
func ReConstructBinaryTree(pre, in []int) *TreeNode {
	if nil == pre || nil == in {
		return nil
	}

	return reConstructBinaryTree(pre, in, 0, len(pre)-1, 0, len(in)-1)
}
func reConstructBinaryTree(pre, in []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	tree := &TreeNode{Value: pre[preStart], Left: nil, Right: nil}
	if preStart == preEnd && inStart == inEnd {
		return tree
	}
	rootIdx := inStart
	for rootIdx < inEnd {
		if pre[preStart] == in[rootIdx] {
			break
		}
		rootIdx++
	}
	leftLength, rightLength := rootIdx-inStart, inEnd-rootIdx
	if leftLength > 0 {
		tree.Left = reConstructBinaryTree(pre, in, preStart+1, preStart+leftLength, inStart, rootIdx-1)
	}
	if rightLength > 0 {
		tree.Right = reConstructBinaryTree(pre, in, preStart+leftLength+1, preEnd, rootIdx+1, inEnd)
	}
	return tree
}

/**
 * 017-树的子结构
 *
 * 输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）
 *
 * 思路1：
 * 1、层次遍历root1, 找到与root2相同的节点(此步骤非递归)
 * 2、找到后判断以此节点为根节点，是否能在root1中找到与root2相同的树结构(此处判断用递归查找)
 */
func IsSubtree(n1, n2 *TreeNode) bool {
	if nil == n1 || nil == n2 {
		return false
	}
	queue := list.New()
	queue.PushBack(n1)

	for queue.Len() > 0 {
		for size, i := queue.Len(), 0; i < size; i++ {
			front := queue.Front()
			cur := front.Value.(*TreeNode)
			queue.Remove(front)

			if cur.Value == n2.Value {
				if IsSameTree(cur, n2) {
					return true
				}
			}
			if nil != cur.Left {
				queue.PushBack(cur.Left)
			}
			if nil != cur.Right {
				queue.PushBack(cur.Right)
			}
		}
	}
	return false
}

func IsSubtree0(n1, n2 *TreeNode) bool {
	if nil == n1 || nil == n2 {
		return false
	}
	return IsSameTree(n1, n2) || IsSubtree0(n1.Left, n2) || IsSubtree0(n1.Right, n2)
}

func IsSameTree(n1, n2 *TreeNode) bool {
	if nil == n1 && nil == n2 {
		return true
	}
	if nil == n1 || nil == n2 {
		return false
	}
	if n1.Value != n2.Value {
		return false
	}
	isLeft := IsSameTree(n1.Left, n2.Left)
	isRight := IsSameTree(n1.Right, n2.Right)
	return isLeft && isRight
}

/**
 * 018-二叉树的镜像
 *
 * 操作给定的二叉树，将其变换为源二叉树的镜像。
 *
 * 输入描述:
 * 二叉树的镜像定义：
 * 源二叉树
 *     	     8
 *     	   /  \
 *     	  6   10
 *     	 / \  / \
 *     	5  7 9 11
 * 镜像二叉树
 *     	     8
 *     	   /  \
 *     	  10   6
 *     	 / \  / \
 *     	11 9 7  5
 * ————————————————
 *
 * 思路1：
 * 1、交换root节点的左右子树
 * 2、递归交换root.left和root.right的左右子树
 *
 */
func MirrorTree0(node *TreeNode) {
	if nil == node || (nil == node.Left && nil == node.Right) {
		return
	}

	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp

	MirrorTree0(node.Left)
	MirrorTree0(node.Right)
}

func MirrorTree1(node *TreeNode) {
	if nil == node || (nil == node.Left && nil == node.Right) {
		return
	}
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		front := queue.Front()
		cur := front.Value.(*TreeNode)
		queue.Remove(front)

		if nil != cur.Left || nil != cur.Right { // swap
			tmp := cur.Left
			cur.Left = cur.Right
			cur.Right = tmp
		}
		if nil != cur.Left {
			queue.PushBack(cur.Left)
		}
		if nil != cur.Right {
			queue.PushBack(cur.Right)
		}
	}
}

/**
 * 022-从上往下打印二叉树
 *
 * 从上往下打印出二叉树的每个节点，同层节点从左至右打印。
 *
 * 思路:
 * 1、利用队列进行层次遍历
 * 2、每次弹出队列中的一个元素，并把左右孩子加入队列即可
 */
func PrintFromTopToBottom(node *TreeNode) {
	if nil == node {
		return
	}
	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		front := queue.Front()
		cur := front.Value.(*TreeNode)
		queue.Remove(front)

		print(cur.Value.(int))
		print(" ")

		if nil != cur.Left {
			queue.PushBack(cur.Left)
		}
		if nil != cur.Right {
			queue.PushBack(cur.Right)
		}
	}
}

/**
 * 023-二叉搜索树的后序遍历序列
 *
 * 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
 * 如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
 *
 * 思路(递归)：
 * 1、后序遍历的特征为 根节点在序列的最后 值为rootVal
 * 2、序列上半部分的值都小于rootVal，下部分的值都大于rootVal
 * 3、递归判断上半部分、下半部分的序列，是否是树的后序遍历序列
 */
func VerifySequenceOfBST(postOrder []int) bool {
	if nil == postOrder || len(postOrder) < 1 {
		return false
	}
	return VerifySequenceOfPostOrderBST(postOrder, 0, len(postOrder)-1)
}
func VerifySequenceOfPostOrderBST(postOrder []int, begin, end int) bool {
	if begin >= end { // 一个元素时，为后序遍历序列
		return true
	}
	// rootVal 节点的值; leftEnd 序列中的最后一个左子树节点;
	rootVal, leftEnd, i := postOrder[end], begin, begin
	for postOrder[i] < rootVal { // 遍历找到左子树的序列 与 右子树序列, 获取分割索引
		leftEnd = i
		i++
	}
	for i < end { // 判断leftEnd序列后的值,如果存在元素小于rootVal,则不是后序序列
		if postOrder[i] < rootVal {
			return false
		}
		i++
	}
	return VerifySequenceOfPostOrderBST(postOrder, begin, leftEnd) && VerifySequenceOfPostOrderBST(postOrder, leftEnd+1, end-1)
}

/**
 * 024-二叉树中和为某一值的路径
 *
 * 题目描述：
 * 输入一颗二叉树的跟节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
 * 路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
 * (注意: 在返回值的list中，数组长度大的数组靠前)
 *
 * 思路：
 *  1、用深度优先搜索DFS
 *  2、每当DFS搜索到新节点时，都要保存该节点。而且每当找出一条路径之后，都将这个保存到list的路径保存到最终结果二维list中
 *  3、并且，每当DFS搜索到子节点，发现不是路径和时，返回上一个结点时，需要把该节点从list中移除
 */
func FindPathInBinaryTree(node *TreeNode, target int) *list.List {
	if nil == node {
		return nil
	}
	result := list.New()
	findPath(result, list.New(), target, node)
	return result
}
func findPath(result, one *list.List, target int, node *TreeNode) {
	if nil == node {
		return
	}
	one.PushBack(node.Value)
	remainVal := target - node.Value.(int)
	if 0 == remainVal && nil == node.Left && nil == node.Right {
		tmp := list.New()
		cur := one.Front()
		for nil != cur {
			tmp.PushBack(cur.Value)
			cur = cur.Next()
		}

		result.PushBack(tmp)
	}
	findPath(result, one, remainVal, node.Left)
	findPath(result, one, remainVal, node.Right)
	one.Remove(one.Back())
}

/**
 * 026-二叉搜索树与双向链表
 *
 * 题目描述
 * 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
 *
 * 思路
 * 1、中序遍历BST的结果是排序的
 * 2、按照遍历顺序组建为双向链表即可
 *
 *
 * 中序遍历过程：
 * 1、利用栈与当前节点的指针
 * 2、处理根节点，如果有左孩子，就处理左孩子，记左孩子为链表的最后一个节点，后续指向当前节点
 * 3、如果左孩子为空，记录当前节点为链表的最后一个节点
 * 4、处理右孩子节点
 */
func ConvertTree2LinkedList(node *TreeNode) *TreeNode {
	if nil == node {
		return nil
	}
	var head, pre *TreeNode
	stack := list.New()
	for cur := node; stack.Len() > 0 || nil != cur; {
		for nil != cur {
			stack.PushFront(cur)
			cur = cur.Left
		}

		front := stack.Front()
		cur = front.Value.(*TreeNode)
		stack.Remove(front)
		if nil == pre {
			head = cur
			pre = head
		} else {
			pre.Right = cur
			cur.Left = pre

			pre = cur
		}
		cur = cur.Right
	}
	return head
}

/**
 * 038-二叉树的深度
 */
func DepthOfBinaryTree(node *TreeNode) int {
	if nil == node {
		return 0
	} else {
		ld := DepthOfBinaryTree(node.Left)
		rd := DepthOfBinaryTree(node.Right)
		if ld > rd {
			return ld + 1
		} else {
			return rd + 1
		}
	}
}

/**
 * 039-平衡二叉树
 */
func IsBalanceBinaryTree(node *TreeNode) bool {
	if nil == node {
		return true
	}
	remain := DepthOfBinaryTree(node.Left) - DepthOfBinaryTree(node.Right)
	if remain > 1 || remain < -1 {
		return false
	} else {
		return IsBalanceBinaryTree(node.Left) && IsBalanceBinaryTree(node.Right)
	}
}

/**
 * 057-二叉树的下一个结点
 *
 * 给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。
 * 注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。
 *
 * 思路：
 * 分析二叉树的下一个节点，一共有以下情况：
 * 1.二叉树为空，则返回空；
 * 2.节点右孩子存在，则设置一个指针从该节点的右孩子出发，一直沿着指向左子结点的指针找到的叶子节点即为下一个节点；
 * 3.节点右孩子不存在，节点是根节点，返回null。
 *   节点不是根节点，如果该节点是其父节点的左孩子，则返回父节点；
 *   否则继续向上遍历其父节点的父节点，重复之前的判断，返回结果。
 *
 */
func GetInOrderNextInBinaryTree(node *TreeNode) *TreeNode {
	if nil == node { // 1.二叉树为空，则返回空；
		return nil
	}
	cur := node
	if nil != cur.Right { // 2.节点右孩子存在，则设置一个指针从该节点的右孩子出发，一直沿着指向左子结点的指针找到的叶子节点即为下一个节点；
		cur = cur.Right
		for nil != cur.Left {
			cur = cur.Left
		}
		return cur
	}

	for nil != cur.Parent { // 3.节点不是根节点，如果该节点是其父节点的左孩子，则返回父节点；
		if cur == cur.Parent.Left {
			return cur.Parent
		}
		cur = cur.Parent //    否则继续向上遍历其父节点的父节点，重复之前的判断，返回结果。
	}
	return nil //    节点右孩子不存在，节点是根节点，返回null。
}

/**
 * 058-对称的二叉树
 *
 * 请实现一个函数，用来判断一颗二叉树是不是对称的。
 * 注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的。
 *
 * 思路1：
 * 1、递归判断两侧的节点是否是对称的
 */
func IsSymmetricalTree(node *TreeNode) bool {
	if nil == node {
		return true
	}
	return isSymmetrical(node.Left, node.Right)
}
func isSymmetrical(left, right *TreeNode) bool {
	if nil == left && nil == right {
		return true
	}
	if nil == left || nil == right {
		return false
	}
	if left.Value != right.Value {
		return false
	}
	return isSymmetrical(left.Left, right.Right) && isSymmetrical(left.Right, right.Left)
}

/**
 * 059-按之字形顺序打印二叉树
 *
 * 题意：按照z字形层次遍历二叉树（以根节点所在层为第1层，则第二层的变量从右边节点开始直到最左边节点，第三层遍历则是从最左边开始到最右边）
 * 思路：z字形层次遍历是对层次遍历加上了一个限制条件（即相邻层，从左到右的遍历顺序相反），因此还是可以采用队列来实现，只不过节点接入队列时需要考虑加入的顺序
 */
func ZigzagLevelOrderPrint(node *TreeNode) {
	if nil == node {
		return
	}
	queue := list.New()
	queue.PushBack(node)

	size, depth := queue.Len(), 0
	for size > 0 {
		for i := 0; i < size; i++ {
			if 0 == depth%2 {
				front := queue.Front()
				cur := front.Value.(*TreeNode)
				queue.Remove(front)

				print(cur.Value.(int))
				print(" ")

				if nil != cur.Left {
					queue.PushBack(cur.Left)
				}
				if nil != cur.Right {
					queue.PushBack(cur.Right)
				}
			} else {
				back := queue.Back()
				cur := back.Value.(*TreeNode)
				queue.Remove(back)

				print(cur.Value.(int))
				print(" ")

				if nil != cur.Right {
					queue.PushFront(cur.Right)
				}
				if nil != cur.Left {
					queue.PushFront(cur.Left)
				}
			}
		}
		size = queue.Len()
		depth++
	}

}

/**
 * 060-把二叉树打印成多行
 * 从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。
 *
 * 思路1：
 * 1、初始化一个队列，初始元素为root
 * 2、遍历元素，每次首先获取当前队列的节点个数，即当前队列的size
 * 3、弹出size次元素，则本次遍历到的均为本层的元素
 * 4、每次弹出元素的同时，把元素的左右孩子加入队列，以便下次遍历
 *
 */
func LevelOrderPrint(node *TreeNode) {
	if nil == node {
		return
	}
	queue := list.New()
	queue.PushBack(node)

	size := queue.Len()
	for size > 0 {
		for i := 0; i < size; i++ {
			front := queue.Front()
			cur := front.Value.(*TreeNode)
			queue.Remove(front)

			print(cur.Value.(int))
			print(" ")

			if nil != cur.Left {
				queue.PushBack(cur.Left)
			}
			if nil != cur.Right {
				queue.PushBack(cur.Right)
			}
		}
		size = queue.Len()
	}
}

/**
 * 062-二叉搜索树的第k个结点
 *
 * 给定一棵二叉搜索树，请找出其中的第k小的结点。
 * 例如， （5，3，7，2，4，6，8）    中，按结点数值大小顺序第三小结点的值为4。
 *
 * 思路
 * 1、二叉搜索树的中序遍历有序递增
 * 2、中序遍历二叉树(递归/非递归)，遍历到第k个元素即停止，则得到第k小元素
 *
 */
func KthInBST(node *TreeNode, k int) *TreeNode {
	if nil == node || k < 1 {
		return nil
	}
	one := list.New()
	helpKthInBST(one, k, node)
	if one.Len() >= k {
		cur := one.Back()
		for i := one.Len() - k; i > 0; i-- {
			cur = cur.Prev()
		}
		return cur.Value.(*TreeNode)
	} else {
		return nil
	}
}
func helpKthInBST(one *list.List, k int, node *TreeNode) {
	if nil == node {
		return
	}
	helpKthInBST(one, k, node.Left)
	one.PushBack(node)
	if one.Len() >= k {
		return
	}
	helpKthInBST(one, k, node.Right)
}
