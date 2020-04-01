package algorithm

import "testing"

func TestFindGreatestSum(t *testing.T) {
	arr := []int{6, -3, -2, 7, -15, 1, 2, 2}
	res := FindGreatestSum(arr)
	println(res)
}
