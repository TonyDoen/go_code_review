package medium

import "testing"

func TestGetMissingNumber(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	res := GetMissingNumber(arr)
	println(res)
}

func TestGetMissingNumberXOR(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	res := GetMissingNumberXOR(arr)
	println(res)
}
