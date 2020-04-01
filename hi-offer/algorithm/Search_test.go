package algorithm

import "testing"

func TestTwoDimensionalArrayLookup(t *testing.T) {
	arr := [][]int{
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{13, 15, 17, 19}}
	res := TwoDimensionalArrayLookup(arr, 15)
	println(res)
}

func TestMoveLeftFindMin(t *testing.T) {
	arr := []int{3, 3, 3, 3, 4, 5, 1, 2, 3, 3}
	res := MoveLeftFindMin(arr)
	println(res)
}

func TestGetTargetCnt(t *testing.T) {
	arr := []int{0, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5}
	target := 3
	res := GetTargetCnt(arr, target)
	println(res)
}
