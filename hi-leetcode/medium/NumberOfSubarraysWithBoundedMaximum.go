package medium

import "math"

/**
 * url: http://www.cnblogs.com/grandyang/p/9237967.html
 * url: https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/discuss/117585/Java-9-liner
 *
 * We are given an array A of positive integers, and two positive integers L and R (L <= R).
 * Return the number of (contiguous*, non-empty) subarrays such that the value of the maximum array element in that subarray is at least L and at most R.
 *
 * Example :
 * Input:
 * A = [2, 1, 4, 3]
 * L = 2
 * R = 3
 * Output: 3
 * Explanation: There are three subarrays that meet the requirements: [2], [2, 1], [3].
 *
 * Note:
 * L, R  and A[i] will be an integer in the range [0, 10^9].
 * The length of A will be in the range of [1, 50000].
 *
 *
 * 题意：
 * 有界限最大值的子数组数量
 *
 * 思路：
 * 既然是求子数组的问题，那么最直接，最暴力的方法就是遍历所有的子数组，然后维护一个当前的最大值，只要这个最大值在[L, R]区间的范围内，结果res自增1即可
 * 这种最原始，最粗犷的暴力搜索法，OJ不答应
 * 优化的方法是，首先，如果当前数字大于R了，那么其实后面就不用再遍历了，不管当前这个数字是不是最大值，它都已经大于R了，那么最大值可能会更大，所以没有必要再继续遍历下去了。同样的剪枝也要加在内层循环中加，当curMax大于R时，直接break掉内层循环即可
 *
 */
func NumberSubarrayBoundedMax0(arr []int, l, r int) int {
	if nil == arr || l > r {
		return -1
	}
	result, length := 0, len(arr)
	for i := 0; i < length; i++ {
		if arr[i] > r {
			continue
		}
		curMax := math.MinInt64
		for j := i; j < length; j++ {
			if curMax < arr[j] {
				curMax = arr[j]
			}
			if curMax > r {
				break
			}
			if curMax >= l {
				result++
			}
		}
	}
	return result
}

func NumberSubarrayBoundedMax1(arr []int, l, r int) int {
	if nil == arr || l > r {
		return -1
	}
	up := helpNumberSubarrayBoundedMax1(arr, r)
	down := helpNumberSubarrayBoundedMax1(arr, l-1)
	return up - down
}
func helpNumberSubarrayBoundedMax1(arr []int, bound int) int {
	result, cur := 0, 0
	for _, x := range arr {
		if x <= bound {
			cur = cur + 1
		} else {
			cur = 0
		}
		result += cur
	}
	return result
}
