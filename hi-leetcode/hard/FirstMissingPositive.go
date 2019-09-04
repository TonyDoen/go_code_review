package hard

import "math"

/**
 * Given an unsorted integer array, find the smallest missing positive integer.
 *
 * Example 1:
 * Input: [1,2,0]
 * Output: 3
 *
 * Example 2:
 * Input: [3,4,-1,1]
 * Output: 2
 *
 * Example 3:
 * Input: [7,8,9,11,12]
 * Output: 1
 * Note:
 *
 * Your algorithm should run in O(n) time and uses constant extra space.
 */

/**
 * 这道题让我们找缺失的首个正数，
 *
 * 由于限定了O(n)的时间，所以一般的排序方法都不能用，最开始我没有看到还限制了空间复杂度，所以想到了用HashSet来解，这个思路很简单，第一遍遍历数组把所有的数都存入HashSet中，并且找出数组的最大值，下次循环从1开始递增找数字，哪个数字找不到就返回哪个数字，如果一直找到了最大的数字，则返回最大值+1，
 */

func GetFirstMissingPositive1(arr []int) int {
	mp := make(map[int]int)
	max := float64(0)
	for _, v := range arr {
		if v < 0 {
			continue
		}
		mp[v] = 1
		max = math.Max(max, float64(v))
	}
	x := int(max)
	for i := 1; i <= x; i++ {
		if i >= 0 && 1 != mp[i] {
			return i
		}
	}
	return x + 1
}

func GetFirstMissingPositiveConstantSpace1(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		v := arr[i]
		vm1 := v - 1
		for v > 0 && v <= n && arr[vm1] != v {
			arr[i] = arr[vm1]
			arr[vm1] = v
		}
	}
	for i := 1; i < n; i++ {
		if i+1 != arr[i] {
			return i + 1
		}
	}
	return n + 1
}
