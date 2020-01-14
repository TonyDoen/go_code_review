package medium

/**
 * Given an array of integers nums, sort the array in ascending order.
 *
 * Example 1:
 * Input: [5,2,3,1]
 * Output: [1,2,3,5]
 *
 * Example 2:
 * Input: [5,1,1,2,0,0]
 * Output: [0,0,1,1,2,5]
 *
 * Note:
 * 1 <= A.length <= 10000
 * -50000 <= A[i] <= 50000
 */

/**
 * 题意：数组排序
 * 思路：记数排序
 * 题目给定了每个数字的范围是 [-50000, 50000]，并不是特别大，这里可以使用记数排序 Count Sort
 * 建立一个大小为 100001 的数组 count，然后统计 arr 中每个数字出现的个数
 * 然后再从0遍历到 100000，对于每个遍历到的数字i，若个数不为0，则加入 count 数组中对应个数的 i-50000 到结果数组中，这里的 50000 是 offset，因为数组下标不能为负数，在开始统计个数的时候，每个数字都加上了 50000，那么最后组成有序数组的时候就要减去，
 */
const NUM = 50000

func CountSort(arr []int) []int {
	count := make([]int, NUM*2+1)
	length := 0
	for _, v := range arr {
		count[v+NUM]++
		length++
	}
	j, res := 0, make([]int, length)
	for i := 0; i < NUM*2+1; i++ {
		for ; count[i] > 0; count[i]-- {
			res[j] = i - NUM
			j++
		}
	}
	return res
}

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}
func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot, i, j := arr[start], start + 1, end
	for i <= j {
		if arr[i] > pivot && arr[j] < pivot {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp

			i++
			j--
		}
		if arr[i] <= pivot {
			i++
		}
		if arr[j] >= pivot {
			j--
		}
	}
	tmp := arr[start]
	arr[start] = arr[j]
	arr[j] = tmp

	quickSort(arr, start, j - 1)
	quickSort(arr, j + 1, end)
}
