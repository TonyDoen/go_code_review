package medium

/**
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
 *
 * For example,
 * Given nums = [0, 1, 3] return 2.
 *
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
 *
 *
 * 题意：
 * 这道题给我们n个数字，是0到n之间的数但是有一个数字去掉了，让我们寻找这个数字，要求线性的时间复杂度和常数级的空间复杂度。
 *
 * 那么最直观的一个方法是用等差数列的求和公式求出0到n之间所有的数字之和，然后再遍历数组算出给定数字的累积和，然后做减法，差值就是丢失的那个数字
 */
func GetMissingNumber(arr []int) int {
	if nil == arr {
		return -1
	}
	sum, length := 0, len(arr)
	for i := 0; i < length; i++ {
		sum += arr[i]
	}
	return length*(length+1)/2 - sum
}

/**
 * 这题还有一种解法，使用位操作Bit Manipulation来解的，用到了异或操作的特性
 * 那么思路是既然0到n之间少了一个数，我们将这个少了一个数的数组合0到n之间完整的数组异或一下，那么相同的数字都变为0了，剩下的就是少了的那个数字了
 */
func GetMissingNumberXOR(arr []int) int {
	if nil == arr {
		return -1
	}
	sum, length := 0, len(arr)
	for i := 0; i < length; i++ {
		sum ^= (i + 1) ^ arr[i]
	}
	return sum
}