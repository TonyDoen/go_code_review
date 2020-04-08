package medium

import "testing"

func TestEventualSafeNodes2(t *testing.T) {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	result := EventualSafeNodes2(graph)
	printList(result)
}
