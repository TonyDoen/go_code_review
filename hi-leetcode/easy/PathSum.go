package easy

import "container/list"

/**
 * url: https://www.cnblogs.com/grandyang/p/4036961.html
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
 * Note: A leaf is a node with no children.
 *
 * Example:
 * Given the below binary tree and sum = 22,
 *
 *       5
 *      / \
 *     4   8
 *    /   / \
 *   11  13  4
 *  /  \      \
 * 7    2      1
 *
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
 */
/**
 * 题意：给了一棵二叉树，问是否存在一条从跟结点到叶结点到路径，使得经过到结点值之和为一个给定的 sum 值
 * 思路：用深度优先算法 DFS 的思想来遍历每一条完整的路径，也就是利用递归不停找子结点的左右子结点，而调用递归函数的参数只有当前结点和 sum 值。
 */
func HasPathSum(root *Node, sum int) bool {
	if nil == root {
		return false
	}
	sum = sum - root.Data
	if nil == root.Left && nil == root.Right && 0 == sum {
		return true
	}
	return HasPathSum(root.Left, sum) || HasPathSum(root.Right, sum)
}

func HasPathSum1(root *Node, sum int) bool {
	if nil == root {
		return false
	}
	stack := list.New()
	stack.PushFront(root)
	for stack.Len() > 0 {
		topEl := stack.Front()
		stack.Remove(topEl)

		top := topEl.Value.(*Node)
		if nil == top.Left && nil == top.Right {
			if sum == top.Data {
				return true
			}
		}
		if nil != top.Right {
			top.Right.Data += top.Data
			stack.PushFront(top.Right)
		}
		if nil != top.Left {
			top.Left.Data += top.Data
			stack.PushFront(top.Left)
		}
	}
	return false
}

/**
 * url: https://www.cnblogs.com/grandyang/p/4036961.html
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
 *
 * For example:
 * Given the below binary tree and sum = 22,
 *
 *       5
 *      / \
 *     4   8
 *    /   / \
 *   11  13  4
 *  /  \    / \
 * 7    2  5   1
 *
 * return
 * [
 *    [5,4,11,2],
 *    [5,8,4,5]
 * ]
 */
/**
 * 题意：给了一棵二叉树，问是否存在一条从跟结点到叶结点到路径，使得经过到结点值之和为一个给定的 sum 值
 * 思路：用深度优先算法 DFS 的思想来遍历每一条完整的路径，也就是利用递归不停找子结点的左右子结点，而调用递归函数的参数只有当前结点和 sum 值。
 */
func GetPathSum(root *Node, sum int) *list.List {
	result := list.New()
	helpGetPathSum(result, list.New(), root, sum)
	return result
}
func helpGetPathSum(result, path *list.List, node *Node, sum int) {
	// base case: leaf
	if nil == node.Left && nil == node.Right {
		if node.Data == sum {
			path.PushBack(node.Data)
			// copy path
			one := list.New()
			tmp := path.Front()
			for nil != tmp {
				one.PushBack(tmp.Value)
				tmp = tmp.Next()
			}
			result.PushBack(one)
			// recover before scenario
			path.Remove(path.Back())
		}
		return
	}

	// recursion rule: non-leaf
	path.PushBack(node.Data)
	if nil != node.Left {
		helpGetPathSum(result, path, node.Left, sum-node.Data)
	}
	if nil != node.Right {
		helpGetPathSum(result, path, node.Right, sum-node.Data)
	}
	path.Remove(path.Back())
}

func GetPathSum1(root *Node, sum int) *list.List {
	result := list.New()
	if nil == root {
		return result
	}

	stack, path, cur, pre := list.New(), list.New(), root, root
	for curSum := 0; nil != cur || stack.Len() > 0; {
		for nil != cur {                                        // 1. 当前节点cur只要不为空，先走到树的最左边节点(第一个while循环)；
			stack.PushFront(cur)
			path.PushBack(cur.Data)

			curSum += cur.Data
			cur = cur.Left
		}
		cur = stack.Front().Value.(*Node)                       // 2. 然后取栈顶元素，但是此时还要继续判断栈顶的右孩子的左子树，
		                                                        //    此时不能pop()，因为有孩子还有可能也是有左子树的；

		if nil != cur.Right && pre != cur.Right {               // 3. pre节点的作用是为了回溯，记录前一个访问的节点，如果cur.right == pre，则说明右子树正在回溯，下面的已经访问完了；
			cur = cur.Right
		} else {                                                // 4. 右孩子为空　或者　已经访问过 此时先判断是否叶子 然后 开始回溯
			if nil == cur.Left && nil == cur.Right && curSum == sum {
				// copy path
				one := list.New()
				tmp := path.Front()
				for nil != tmp {
					one.PushBack(tmp.Value)
					tmp = tmp.Next()
				}
				result.PushBack(one)
			}
			stack.Remove(stack.Front()) // pop
			path.Remove(path.Back())    // remove last

			pre = cur
			curSum -= cur.Data
			cur = nil
		}
	}
	return result
}
