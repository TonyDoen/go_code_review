package medium

import "math"

/**
 * Given a circular array C of integers represented by A, find the maximum possible sum of a non-empty subarray of C.
 * Here, a circular array means the end of the array connects to the beginning of the array.  (Formally, C[i] = A[i] when 0 <= i < A.length, and C[i+A.length] = C[i] when i >= 0.)
 * Also, a subarray may only include each element of the fixed buffer A at most once.  (Formally, for a subarray C[i], C[i+1], ..., C[j], there does not exist i <= k1, k2 <= j with k1 % A.length = k2 % A.length.)
 *
 * Example 1:
 * Input: [1,-2,3,-2]
 * Output: 3
 * Explanation: Subarray [3] has maximum sum 3
 *
 * Example 2:
 * Input: [5,-3,5]
 * Output: 10 Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10
 *
 * Example 3:
 * Input: [3,-1,2,-1]
 * Output: 4
 * Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4
 *
 * Example 4:
 * Input: [3,-2,2,-3]
 * Output: 3 Explanation: Subarray [3] and [3,-2,2] both have maximum sum 3
 *
 * Example 5:
 * Input: [-2,-3,-1]
 * Output: -1 Explanation: Subarray [-1] has maximum sum -1
 *
 * Note:
 * -30000 <= A[i] <= 30000
 * 1 <= A.length <= 30000
 */

/**
 * 题意：
 * 这道题让求环形子数组的最大和，对于环形数组，我们应该并不陌生，之前也做过类似的题目 Circular Array Loop，就是说遍历到末尾之后又能回到开头继续遍历。假如没有环形数组这一个条件，其实就跟之前那道 Maximum Subarray 一样，解法比较直接易懂。这里加上了环形数组的条件，难度就增加了一些，需要用到一些 trick。既然是子数组，则意味着必须是相连的数字，而由于环形数组的存在，说明可以首尾相连，这样的话，最长子数组的范围可以有两种情况，一种是正常的，数组中的某一段子数组，另一种是分为两段的，即首尾相连的，可以参见 大神 lee215 的帖子 中的示意图。对于第一种情况，其实就是之前那道题 Maximum Subarray 的做法，对于第二种情况，需要转换一下思路，除去两段的部分，中间剩的那段子数组其实是和最小的子数组，只要用之前的方法求出子数组的最小和，用数组总数字和一减，同样可以得到最大和。两种情况的最大和都要计算出来，取二者之间的较大值才是真正的和最大的子数组。但是这里有个 corner case 需要注意一下，假如数组中全是负数，那么和最小的子数组就是原数组本身，则求出的差值是0，而第一种情况求出的和最大的子数组也应该是负数，那么二者一比较，返回0就不对了，所以这种特殊情况需要单独处理一下，
 */
func maxSubarraySumCircular(arr []int) int {
	var sum, mn, mx, curMax, curMin = 0, math.MaxInt32, math.MinInt32, 0, 0
	for _, v := range arr {
		curMin = min(curMin+v, v)
		mn = min(mn, curMin)
		curMax = max(curMax+v, v)
		mx = max(mx, curMax)
		sum += v
	}
	if sum-mn == 0 {
		return mx
	} else {
		return max(mx, sum-mn)
	}
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
