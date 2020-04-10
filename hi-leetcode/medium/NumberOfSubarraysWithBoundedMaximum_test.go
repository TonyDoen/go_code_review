package medium

import "testing"

func TestNumberSubarrayBoundedMax0(t *testing.T) {
	arr, l, r := []int{2, 1, 4, 3}, 2, 3
	res := NumberSubarrayBoundedMax0(arr, l, r)
	println(res)
}

func TestNumberSubarrayBoundedMax1(t *testing.T) {
	arr, l, r := []int{2, 1, 4, 3}, 2, 3
	res := NumberSubarrayBoundedMax1(arr, l, r)
	println(res)
}
