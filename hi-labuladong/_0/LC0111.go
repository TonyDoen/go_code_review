package _0

import (
	"container/list"
)

/**
 * BFS
 *
 * int bfs(Node start, Node target) {
 *     Queue<Node> q;     //
 *     Set<Node> visited; // 避免走回头路
 *
 *     q.offer(start);    // 将起点加入队列
 *     visited.add(start);
 *     int step = 0;      // 记录扩散的步数
 *
 *     while (q.notEmpty()) {
 *         int sz = q.size();
 *
 *         // 将当前队列中的所有节点向四周扩散
 *         for (int i = 0; i < sz; i++) {
 *             Node cur = q.poll();
 *             // 重点： 判断是否到达终点
 *             if (cur.equals(target)) {
 *                 return step;
 *             }
 *
 *             // 将 cur 相邻节点加入队列
 *             for (Node x : cur.adj()) {
 *                 if (!visited.contains(x)) {
 *                     q.offer(x);
 *                     visited.add(x);
 *                 }
 *             }
 *         }
 *
 *         // 重点： 更新布数
 *         step++;
 *     }
 * }
 */

/**
 * LeetCode 第 111 题: 判断一棵二叉树的最小高度
 *
 * 思路：
 * 1. 显然起点就是 root 根节点,终点就是最靠近根节点的那个「叶子节点」
 * 2. 叶子节点就是两个子节点都是 null 的节点 => if (cur.left == null && cur.right == null) // 到达叶子节点
 */
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func MinDepth(node *TreeNode) int {
	if nil == node {
		return 0
	}

	var q = list.New()
	q.PushBack(node)
	var depth = 1 // root node is first level, then depth = 1

	for q.Len() > 0 {
		var sz = q.Len()
		// expand from this node
		for i := 0; i < sz; i++ {
			// pop poll
			var front = q.Front()
			var cur = front.Value.(*TreeNode)
			q.Remove(front)

			// careful
			if nil == cur.Left && nil == cur.Right {
				return depth
			}

			// add [cur]'s neighbors
			if nil != cur.Left {
				q.PushBack(cur.Left)
			}
			if nil != cur.Right {
				q.PushBack(cur.Right)
			}
		}

		// careful
		depth++
	}
	return depth
}
