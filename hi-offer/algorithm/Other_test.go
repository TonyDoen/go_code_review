package algorithm

import "testing"

// 002-替换空格
func TestReplaceBlank20(t *testing.T) {
	src := "We Are Happy."
	res := ReplaceBlank20(src)
	println(res)

	res = ReplaceBlank201(src)
	println(res)
}

// 013-调整数组顺序使奇数位于偶数前面
func TestOddBeforeEven(t *testing.T) {
	arr := []int{0, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5}
	OddBeforeEven(arr)
	for _, v := range arr {
		print(v)
		print(" ")
	}
}

func TestOddBeforeEven2(t *testing.T) {
	arr := []int{0, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5}
	OddBeforeEven2(arr)
	for _, v := range arr {
		print(v)
		print(" ")
	}
}

// 028-数组中出现次数超过一半的数字
func TestFindMoreThanHalfNumber(t *testing.T) {
	arr := []int{0, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5}
	res := FindMoreThanHalfNumber(arr)
	println(res)
}

func TestOddQuickSort(t *testing.T) {
	arr := []int{3, 32, 321}
	OddQuickSort(arr, 0, len(arr)-1)
	for _, v := range arr {
		print(v)
		print(" ")
	}
}

// 032-把数组排成最小的数
func TestPrintOddQuickSortNumber(t *testing.T) {
	arr := []int{3, 32, 321}
	PrintOddQuickSortNumber(arr)
}

// 033-丑数
func TestGetUglyNumber(t *testing.T) {
	n := 10
	res := GetUglyNumber(n)
	println(res)
}