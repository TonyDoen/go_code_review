package algorithm

import "container/list"

/**
 * 全排列
 * 027-字符串的排列
 *
 * 题目描述：
 * 输入一个字符串,按字典序打印出该字符串中字符的所有排列。
 * 例如输入字符串abc,则打印出由字符a,b,c 所能排列出来的所有字符串
 * abc,acb,bac,bca,cab和cba。
 *
 * 输入描述:
 * 输入一个字符串,长度不超过9(可能有字符重复),字符只包括大小写字母。
 */
func Permutation(src string) *list.List {
	if len(src) < 1 {
		return nil
	}

	result := list.New()
	helpPermutation(result, 0, []rune(src))
	return result
}
func helpPermutation(result *list.List, begin int, arr []rune) {
	length := len(arr)
	if length-1 == begin {
		result.PushBack(string(arr))
	}
	for i := begin; i < length; i++ {
		// 与begin不同位置的元素相等，不需要交换
		if i != begin && arr[i] == arr[begin] {
			continue
		}

		swap(arr, begin, i)                   // 交换元素
		helpPermutation(result, begin+1, arr) // 处理后续元素
		swap(arr, begin, i)                   // 数组复原
	}
}
func swap(arr []rune, i, j int) {
	if i == j {
		return
	}
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

/**
 * 035-数组中的逆序对(归并排序)(冒泡排序)
 *
 * 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
 * 输入一个数组,求出这个数组中的逆序对的总数P。
 * 并将P对1000000007取模的结果输出。 即输出P%1000000007
 *
 * 1. 冒泡排序：
 * 交换次数就是逆序对数，换句换说就是交换一次的两个数就是一组逆序对。
 * 时间复杂度 O(n^2) ；
 * 空间复杂度 O(1)。
 *
 * 2. 归并排序
 * 归并排序分而治之，统计逆序对数；两数组第一个数比较后数组小的时候，逆序对数加上前数组的长度，其他与归并排序完全一样。
 * 时间复杂度 O(N*logN)；
 * 空间复杂度 O(n).
 *
 *            [ 7, 5, 6, 4 ]            1> 把长度 4 的数组 拆分成 长度为 2 的数组
 *          [ 7, 5 ]; [ 6, 4 ]          2> 把长度 2 的数组 拆分成 长度为 1 的数组
 *      [ 7 ]; [ 5 ]; [ 6 ]; [ 4 ];     3> 把长度 1 的数组 合并，排序，统计逆序对
 *          [ 5, 7 ]; [ 4, 6 ]          4> 把长度 2 的数组 合并，排序，统计逆序对
 *           [ 4, 5, 6, 7 ]
 */
func InversePair(arr []int) int {
	if nil == arr || len(arr) < 1 {
		return -1
	}
	result := 0
	mergeSort(&result, 0, len(arr)-1, arr)
	return result
}
func mergeSort(cnt *int, left, right int, arr []int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(cnt, left, mid, arr)
	mergeSort(cnt, mid+1, right, arr)

	// merge
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else { // 左比右大，左之后所有的都比右大，有mid-i+1个逆序对
			tmp[k] = arr[j]
			k++
			j++
			*cnt += mid - i + 1
			*cnt %= 1000000007 // 防止溢出，每一步都要取余
		}
	}
	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= right {
		tmp[k] = arr[j]
		k++
		j++
	}
	tmpLength := len(tmp)
	for l := 0; l < tmpLength; l++ {
		arr[left+l] = tmp[l]
	}
}
