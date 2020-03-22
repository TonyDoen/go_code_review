package medium

import "container/list"

/**
 * Clone Graph
 * Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph. Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.
 *
 * 1 - 2
 * |   |
 * 4 - 3
 *
 * Input:
 * {"$id":"1","neighbors":[{"$id":"2","neighbors":[{"$ref":"1"},{"$id":"3","neighbors":[{"$ref":"2"},{"$id":"4","neighbors":[{"$ref":"3"},{"$ref":"1"}],"val":4}],"val":3}],"val":2},{"$ref":"4"}],"val":1}
 *
 * Explanation:
 * Node 1's value is 1, and it has two neighbors: Node 2 and 4.
 * Node 2's value is 2, and it has two neighbors: Node 1 and 3.
 * Node 3's value is 3, and it has two neighbors: Node 2 and 4.
 * Node 4's value is 4, and it has two neighbors: Node 1 and 3.
 *
 * Note:
 * The number of nodes will be between 1 and 100.
 * The undirected graph is a simple graph, which means no repeated edges and no self-loops in the graph.
 * Since the graph is undirected, if node p has node q as neighbor, then node q must have node p as neighbor too.
 * You must return the copy of the given node as a reference to the cloned graph.
 */
/**
 *  克隆无向图
 * 1、递归实现
 * 2、具体实现为：复制原节点，挂接邻居节点的复制节点到原节点的复制节点下
 * 3、怎么复制邻居节点呢？递归调用复制原节点的函数
 * 4、递归出口： 节点为null || 已经复制过该节点了
 */

type UndirectedGraphNode struct {
	Label     int
	Neighbors *list.List
}

func (node *UndirectedGraphNode)CloneGraph() *UndirectedGraphNode {
	if nil == node {
		return nil
	}
	cache := make(map[*UndirectedGraphNode]*UndirectedGraphNode)
	return cloneGraph(node, cache)
}

func cloneGraph(node *UndirectedGraphNode, cache map[*UndirectedGraphNode]*UndirectedGraphNode) *UndirectedGraphNode {
	if nil == node {
		return nil
	}
	val := cache[node]
	if nil != val {
		return val
	}

	// copyNode label
	copyNode := &UndirectedGraphNode{Label: node.Label, Neighbors: list.New()}
	cache[node] = copyNode
	// copyNode neighbors
	cur := node.Neighbors.Front()
	for nil != cur {
		if nil != cur.Value {
			val := cur.Value.(*UndirectedGraphNode)
			one := cloneGraph(val, cache)
			copyNode.Neighbors.PushBack(one)
		}
		cur = cur.Next()
	}

	return copyNode
}
