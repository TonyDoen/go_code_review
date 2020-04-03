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

// 041-和为S的连续正数序列(滑动窗口思想)
func TestFindContinuousSequence(t *testing.T) {
	sum := 100
	res := FindContinuousSequence(sum)
	cur := res.Front()
	for nil != cur {
		arr := cur.Value.([]int)
		for _, v := range arr {
			print(v)
			print(" ")
		}
		cur = cur.Next()
		println()
	}
}

// 042-和为S的两个数字(双指针思想)
func TestFindNumberWithSum(t *testing.T) {
	arr := []int{-2, -1, 0, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5, 7, 8}
	s := 6
	num1, num2 := FindNumberWithSum(arr, s)
	println(num1, ",", num2)
}

// 043-左旋转字符串(矩阵翻转)
func TestLeftRotateString(t *testing.T) {
	src := "abcXYZdef"
	n := 3
	res := LeftRotateString(src, n)
	println(res)
}

// 046-孩子们的游戏-圆圈中最后剩下的数(约瑟夫环)
func TestJosephusCircle(t *testing.T) {
	n, m := 3, 2
	res := JosephusCircle(n, m)
	println(res)
}

// 051-构建乘积数组
func TestProductOfArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	res := ProductOfArray(arr)
	for _, v := range res {
		print(v)
		print(" ")
	}
}
