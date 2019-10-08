package hard

import (
	"fmt"
	"testing"
)

func TestReachableNodes(t *testing.T) {
	edges := [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}
	m := 6
	n := 3
	res := ReachableNodes(edges, m, n)
	fmt.Println(res)
}
