package medium

/**
 * url:http://www.cnblogs.com/grandyang/p/8519566.html
 * url:https://leetcode.com/problems/is-graph-bipartite/discuss/115487/Java-Clean-DFS-solution-with-Explanation
 *
 * Given an undirected graph, return true if and only if it is bipartite.
 * Recall that a graph is bipartite if we can split it's set of nodes into two independent subsets A and B such that every edge in the graph has one node in A and another node in B.
 * The graph is given in the following form: graph[i] is a list of indexes j for which the edge between nodes i and j exists.  Each node is an integer between 0 and graph.length - 1.  There are no self edges or parallel edges: graph[i] does not contain i, and it doesn't contain any element twice.
 *
 * Example 1:
 * Input: [[1,3], [0,2], [1,3], [0,2]]
 * Output: true
 *
 * Explanation:
 * The graph looks like this:
 * 0----1
 * |    |
 * |    |
 * 3----2
 * We can divide the vertices into two groups: {0, 2} and {1, 3}.
 *
 *
 * Example 2:
 * Input: [[1,2,3], [0,2], [0,1,3], [0,2]]
 * Output: false
 *
 * Explanation:
 * The graph looks like this:
 * 0----1
 * | \  |
 * |  \ |
 * 3----2
 * We cannot find a way to divide the set of nodes into two independent subsets.
 *
 * Note:
 * graph will have length in range [1, 100].
 * graph[i] will contain integers in range [0, graph.length - 1].
 * graph[i] will not contain i or duplicate values.
 * The graph is undirected: if any element j is in graph[i], then i will be in graph[j].
 *
 *
 * 题意: 是二分图吗
 * 输入数组中的graph[i]，表示顶点i所有相邻的顶点，比如对于例子1来说，顶点0和顶点1，3相连，顶点1和顶点0，2相连，顶点2和结点1，3相连，顶点3和顶点0，2相连。
 * 这道题让我们验证给定的图是否是二分图，所谓二分图，就是可以将图中的所有顶点分成两个不相交的集合，使得同一个集合的顶点不相连。为了验证是否有这样的两个不相交的集合存在，
 * 我们采用一种很机智的染色法，大体上的思路是要将相连的两个顶点染成不同的颜色，一旦在染的过程中发现有两连的两个顶点已经被染成相同的颜色，说明不是二分图。这里我们使用两种颜色，
 * 分别用1和-1来表示，初始时每个顶点用0表示未染色，然后遍历每一个顶点，如果该顶点未被访问过，则调用递归函数，如果返回false，那么说明不是二分图，则直接返回false。如果循环退出后没有返回false，则返回true。
 * 在递归函数中，如果当前顶点已经染色，如果该顶点的颜色和将要染的颜色相同，则返回true，否则返回false。如果没被染色，则将当前顶点染色，然后再遍历与该顶点相连的所有的顶点，调用递归函数，如果返回false了，则当前递归函数的返回false，循环结束返回true，参见代码如下：
 */
func IsBipartite1(graph [][]int) bool {
	if nil == graph {
		return false
	}
	length := len(graph)
	for i, colors := 0, make([]int, length); i < length; i++ {
		if 0 == colors[i] && !validBipartite1(i, 1, colors, graph) {
			return false
		}
	}
	return true
}
func validBipartite1(cur, color int, colors []int, graph [][]int) bool {
	if 0 != colors[cur] {
		return colors[cur] == color
	}
	colors[cur] = color
	for _, i := range graph[cur] {
		if !validBipartite1(i, -1*color, colors, graph) {
			return false
		}
	}
	return true
}

//// 非递归写法
//func IsBipartite2(graph [][]int) bool { // error
//	if nil == graph {
//		return false
//	}
//	length := len(graph)
//	for i, colors := 0, make([]int, length); i < length; i++ {
//		if 0 != colors[i] {
//			continue
//		}
//		colors[i] = 1
//		queue := list.New()
//		queue.PushBack(i)
//		for queue.Len() > 0 {
//			front := queue.Front()
//			t := front.Value.(int)
//			queue.Remove(front)
//
//			for a := range graph[t] {
//				if 0 == colors[a] {
//					colors[a] = -1 * colors[t]
//					queue.PushFront(a)
//				} else {
//					if colors[a] == colors[t] {
//						return false
//					}
//				}
//			}
//		}
//	}
//	return true
//}
