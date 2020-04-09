package medium

import "testing"

func TestIsBipartite1(t *testing.T) {
	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	res := IsBipartite1(graph)
	println(res)
}

//func TestIsBipartite2(t *testing.T) {
//	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
//	res := IsBipartite2(graph)
//	println(res)
//}
