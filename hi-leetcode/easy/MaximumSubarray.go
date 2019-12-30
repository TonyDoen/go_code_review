package easy

import "math"

/**
 * Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
 *
 * Example:
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 * Follow up:
 * If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 */

/**
 * 题意：最大子数组
 * 思路：这道题让求最大子数组之和，并且要用两种方法来解，分别是 O(n) 的解法，还有用分治法 Divide and Conquer Approach，这个解法的时间复杂度是 O(nlgn)，那就先来看 O(n) 的解法，定义两个变量 res 和 curSum，其中 res 保存最终要返回的结果，即最大的子数组之和，curSum 初始值为0，每遍历一个数字 num，比较 curSum + num 和 num 中的较大值存入 curSum，然后再把 res 和 curSum 中的较大值存入 res，以此类推直到遍历完整个数组，可得到最大子数组的值存在 res 中
 */

func MaxSubArray1(arr []int) int {
	var res, curSum = math.MinInt32, 0

	for _, num := range arr {
		curSum = max(curSum + num, num)
		res = max(res, curSum)
	}
	return res
}

/**
 * 思路：题目还要求我们用分治法 Divide and Conquer Approach 来解，这个分治法的思想就类似于二分搜索法，需要把数组一分为二，分别找出左边和右边的最大子数组之和，然后还要从中间开始向左右分别扫描，求出的最大值分别和左右两边得出的最大值相比较取最大的那一个，代码如下：
 */
func MaxSubArray2(arr []int) int {
	length := len(arr)
	if 0 == length {
		return 0
	}
	return helpMaxSubArray2(arr, 0, length - 1)
}
func helpMaxSubArray2(arr []int, left, right int) int {
	if left >= right {
		return arr[left]
	}
	mid := left + (right - left) / 2
	lMax := helpMaxSubArray2(arr, left, mid - 1)
	rMax := helpMaxSubArray2(arr, mid + 1, right)

	mMax := arr[mid]
	t := mMax
	for i := mid-1; i >= left; i-- {
		t += arr[i]
		mMax = max(mMax, t)
	}
	t = mMax
	for i := mid+1; i <= right; i++ {
		t += arr[i]
		mMax = max(mMax, t)
	}
	return max(mMax, max(lMax, rMax))
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
