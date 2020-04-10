package medium

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	arr := []int{1, 3, 5, 4, 7}
	res := FindNumberOfLIS(arr)
	println(res)
}
