package medium

import "testing"

func TestLongestIncreasingPath0(t *testing.T) {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1}}
	res := LongestIncreasingPath0(matrix)
	println(res)
}
