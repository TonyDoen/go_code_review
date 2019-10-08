package easy

import "github.com/TonyDoen/go_code_review/hi-leetcode/medium"

/**
 * An array is monotonic if it is either monotone increasing or monotone decreasing.
 * An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
 * Return true if and only if the given array A is monotonic.
 *
 * Example 1:
 * Input: [1,2,2,3]
 * Output: true
 *
 * Example 2:
 * Input: [6,5,4,4]
 * Output: true
 *
 * Example 3:
 * Input: [1,3,2]
 * Output: false
 *
 * Example 4:
 * Input: [1,2,4,5]
 * Output: true
 *
 * Example 5:
 * Input: [1,1,1]
 * Output: true
 *
 * Note:
 * 1 <= A.length <= 50000
 * -100000 <= A[i] <= 100000
 */

/**
 * 单调数组
 * 这道题让我们判断一个数组是否单调，单调数组就是说这个数组的数字要么是递增的，要么是递减的，不存在一会儿递增一会儿递减的情况，即不会有山峰存在。这里不是严格的递增或递减，是允许有相同的数字的。
 *
 * 思路：
 * 那么我们直接将相邻的两个数字比较一下即可，使用两个标识符，inc 和 dec，初始化均为 true，我们开始时假设这个数组既是递增的又是递减的，当然这是不可能的，我们会在后面对其进行更新。在遍历数组的时候，只要发现某个数字大于其身后的数字了，那么 inc 就会赋值为 false，同理，只要某个数字小于其身后的数字了，dec 就会被赋值为 false，所以在既有递增又有递减的数组中，inc 和 dec 都会变为 false，而在单调数组中二者之间至少有一个还会保持为 true，参见代码如下：
 */

func IsMonotonic(arr []int) bool {
	inc := true
	dec := true
	length := len(arr)
	for i := 1; i < length; i++ {
		inc = inc && (arr[i-1] <= arr[i])
		dec = dec && (arr[i-1] >= arr[i])
		if !inc && !dec {
			return false
		}
	}
	return true
}

func IsMonotonic2(arr []int) bool {
	inc := 1
	dec := 1
	length := len(arr)
	for i := 1; i < length; i++ {
		inc += medium.When(arr[i-1] <= arr[i], 1, 0).(int)
		dec += medium.When(arr[i-1] >= arr[i], 1, 0).(int)
	}
	return (length == inc) || (length == dec)
}
