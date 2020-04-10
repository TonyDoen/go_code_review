package medium

import "testing"

func TestMatrixScore(t *testing.T) {
	src := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	res := MatrixScore(src)
	println(res)
}
