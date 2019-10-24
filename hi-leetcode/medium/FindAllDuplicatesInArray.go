package medium

import (
	"container/list"
	"math"
)

/**
 * Find All Duplicates in an Array
 * Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
 * Find all the elements that appear twice in this array.
 * Could you do it without extra space and in O(n) runtime?
 *
 * Example:
 * Input:
 * [4,3,2,7,8,2,3,1]
 * Output:
 * [2,3]
 *
 * 题意：找出数组中所有重复项
 * 这道题给了我们一个数组，数组中的数字可能出现一次或两次，让我们找出所有出现两次的数字，由于之前做过一道类似的题目Find the Duplicate Number，所以不是完全无从下手。这类问题的一个重要条件就是1 ≤ a[i] ≤ n (n = size of array)，不然很难在O(1)空间和O(n)时间内完成。首先来看一种正负替换的方法，这类问题的核心是就是找nums[i]和nums[nums[i] - 1]的关系，我们的做法是，对于每个nums[i]，我们将其对应的nums[nums[i] - 1]取相反数，如果其已经是负数了，说明之前存在过，我们将其加入结果res中即可，参见代码如下
 */
func findDuplicates(nums []int) *list.List {
	res, length := list.New(), len(nums)
	for i := 0; i < length; i++ {
		idx := int(math.Abs(float64(nums[i]))) - 1
		if nums[idx] < 0 {
			res.PushBack(idx + 1)
		}
		nums[idx] = -nums[idx]
	}
	return res
}

/**
 * 思路：
 * 下面这种方法是在nums[nums[i]-1]位置累加数组长度n，注意nums[i]-1有可能越界，所以我们需要对n取余，最后要找出现两次的数只需要看nums[i]的值是否大于2n即可，最后遍历完nums[i]数组为[12, 19, 18, 15, 8, 2, 11, 9]，我们发现有两个数字19和18大于2n，那么就可以通过i+1来得到正确的结果2和3了，参见代码如下
 */
func findDuplicates3(nums []int) *list.List {
	res, n := list.New(), len(nums)
	for i := 0; i < n; i++ {
		nums[(nums[i]-1)%n] += n
	}
	for i := 0; i < n; i++ {
		if nums[i] > 2*n {
			res.PushBack(i + 1)
		}
	}
	return res
}
