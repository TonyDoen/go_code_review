package algorithm

import "container/list"

/**
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
