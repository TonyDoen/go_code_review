package easy

/**
 * url: https://leetcode.com/problems/sort-array-by-parity/
 * url: https://www.cnblogs.com/grandyang/p/11173513.html
 *
 * Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
 * You may return any answer array that satisfies this condition.
 *
 * Example 1:
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 *
 * Note:
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 */

/**
 * 题意：按奇偶排序数组
 * 这道题让我们给数组重新排序，使得偶数都排在奇数前面，
 *
 * 思路：
 * 并不难。最直接的做法就是分别把偶数和奇数分别放到两个数组中，然后把奇数数组放在偶数数组之后，将拼接成的新数组直接返回即可
 */

func SortArrayByParity1(arr []int) []int {
	length := len(arr)
	even := make([]int, length)
	odd := make([]int, length)
	var ei, oi int
	for _, n := range arr {
		if 0 == n%2 {
			even[ei] = n
			ei++
		} else {
			odd[oi] = n
			oi++
		}
	}
	for i := 0; i < oi; i++ {
		even[ei] = odd[i]
		ei++
	}
	return even
}

func SortArrayByParity2(arr []int) []int {
	length := len(arr)
	i := 0
	for j := 0; j < length; j++ {
		tmp := arr[j]
		if 0 == tmp%2 {
			// swap
			arr[j] = arr[i]
			arr[i] = tmp
			i++
		}
	}
	return arr
}
