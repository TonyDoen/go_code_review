package medium

import (
	"container/list"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	node := prepareUndirectedGraph()

	copyNode := node.CloneGraph()
	println(copyNode)
}

func prepareUndirectedGraph() *UndirectedGraphNode {
	/**
	 * 1 - 2
	 * |   |
	 * 4 - 3
	 */
	_1 := &UndirectedGraphNode{Label: 1, Neighbors: list.New()}
	_2 := &UndirectedGraphNode{Label: 2, Neighbors: list.New()}
	_3 := &UndirectedGraphNode{Label: 3, Neighbors: list.New()}
	_4 := &UndirectedGraphNode{Label: 4, Neighbors: list.New()}

	_1.Neighbors.PushBack(_2)
	_1.Neighbors.PushBack(_4)
	_2.Neighbors.PushBack(_1)
	_2.Neighbors.PushBack(_3)
	_3.Neighbors.PushBack(_2)
	_3.Neighbors.PushBack(_4)
	_4.Neighbors.PushBack(_1)
	_4.Neighbors.PushBack(_3)

	return _1
}
