package medium

import "container/list"

/**
 * All Paths From Source to Target
 *
 * url: https://leetcode.com/problems/all-paths-from-source-to-target/solution/
 * url: http://www.cnblogs.com/grandyang/p/9262159.html
 *
 * Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.
 * The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j for which the edge (i, j) exists.
 *
 * Example:
 * Input: [[1,2], [3], [3], []]
 * Output: [[0,1,3],[0,2,3]]
 * Explanation: The graph looks like this:
 * 0--->1
 * |    |
 * v    v
 * 2--->3
 * There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
 *
 * Note:
 * The number of nodes in the graph will be in the range [2, 15].
 * You can print different paths in any order, but you should keep the order of nodes inside one path.
 *
 *
 * 题意：从起点到目标点到所有路径
 * 给一个有向无环图，二维数组代表图中存在等路径，例如：graph[i]=[j,k] 表示(i,j),(i,k)这两条路径存在。 求从起点到目标点到所有路径
 *
 * 思路：
 * 这道题的本质就是遍历邻接链表，由于要列出所有路径情况，那么递归就是不二之选了。
 * 我们用cur来表示当前遍历到的结点，初始化为0，然后在递归函数中，先将其加入路径path，如果cur等于N-1了，那么说明到达结点N-1了，将path加入结果res。否则我们再遍历cur的邻接结点，调用递归函数即可
 */
func AllPathsFromSource2Target0(graph [][]int) *list.List {
	if nil == graph {
		return nil
	}
	return helpAllPathsFromSource2Target0(0, graph)
}
func helpAllPathsFromSource2Target0(cur int, graph [][]int) *list.List {
	res := list.New()
	if len(graph)-1 == cur { // 1. 如果cur等于N-1了，直接将cur先装入数组，再装入结果res中返回。
		l := list.New()
		l.PushBack(cur)
		res.PushBack(l)
		return res
	}

	for _, neighbor := range graph[cur] { // 2. 否则就遍历cur的邻接结点
		lt := helpAllPathsFromSource2Target0(neighbor, graph) // 2.1 对于每个邻接结点，先调用递归函数，然后遍历其返回的结果
		front := lt.Front()                                   // 2.2 对于每个遍历到的path，将cur加到数组首位置，然后将path加入结果res中
		for nil != front {
			path := front.Value.(*list.List)
			path.PushFront(cur)
			res.PushBack(path)
			front = front.Next()
		}
	}
	return res
}

