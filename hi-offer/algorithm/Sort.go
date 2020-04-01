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

/**
 * 029-最小的K个数(堆排序)(快速排序)
 *
 * 输入n个整数，找出其中最小的K个数。
 * 例如输入 4,5,1,6,2,7,3,8 这8个数字，则最小的4个数字是 1,2,3,4 。
 *
 * 思路1:
 * 按照升序排序，然后取前K个数，就是我们最终想要的到的结果，
 * 现在较好一点的排序方法时间复杂度是N*logN
 *
 */
/**
 * 思路2:
 * 根据一次快排(Partition)的想法，可知一次随机快速排序可以确定一个有序的位置，这个位置的左边都小于这个数，右边都大于这个数，
 * 我们如果能找到随机快速排序确定的位置等于k-1的那个位置，那么0~k-1个数就是我们要找的数。
 * 怎么能找到那个位置：
 * - 如果Partition确定的位置小于K-1，说明k-1这个位置在它的右边，我们继续在右边进行查找。
 * - 如果Partition确定的位置大于K-1，说明k-1这个位置在它的左边，我们继续在左边进行查找。
 *
 * 缺点： 时间复杂度虽然是O(n)，但是找出来的最小的K个数却不是排序过的。而且这种方法有个限制，就是必须修改给的数组。
 *
 */
func FindKthNumber0(arr []int, k int) []int {
	if nil == arr || k < 1 {
		return nil
	}
	start, end, _1k := 0, len(arr)-1, k-1
	pivot := quickPartition(arr, start, end)
	for _1k != pivot {
		if _1k > pivot {
			start = pivot + 1
			pivot = quickPartition(arr, start, end)
		} else if _1k < pivot {
			end = pivot - 1
			pivot = quickPartition(arr, start, end)
		}
	}
	return arr[:pivot+1]
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := quickPartition(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}

func quickPartition(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

/**
 *
 * 思路3:
 * 是建一个K个数的大顶堆，每次拿一个数和堆顶元素比较，如果这个数比堆顶元素大，则必然不是最小的K个数，如果这个数比堆顶元素小，则与堆顶元素交换，然后在向下调整一次建成新的大堆，然后遍历所有的数，直到最小的K个数都进堆里。
 * 最大的K个数 ---- 建小顶堆
 * 最小的K个数 ---- 建大顶堆
 * 优点：海量数据不占内存；实时判别新产生的数据；时间复杂度O(N*logN)。
 *
 *
 * 堆元素下标从0开始，习惯问题，此时除根节点外的任意节点的索引位置关系为
 * 父节点位置：ceiling(i/2) - 1; 左孩子：i*2 + 1; 右孩子：i*2 + 2; 最后一个分支节点位置为：n/2 - 1;
 */
func FindKthNumber1(arr[]int, k int) []int {
	if nil == arr || k < 1 {
		return nil
	}
	length := len(arr)
	if k >= length { // 全排序
		HeapSort(arr)
		return arr
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = arr[i]
	}

	_1k := k-1
	for i := (k-2)/2; i >= 0; i-- { // 构建堆
		downAdjust(result, i, _1k)
	}
	for i := k; i < length; i++ {   // 进行堆排序
		if arr[i] < result[0] {
			result[0] = arr[i]
			downAdjust(result, 0, _1k)
		}
	}
	return result
}

func HeapSort(arr []int) {
	if nil == arr {
		return
	}
	// 构建堆
	length := len(arr)
	for i := (length - 2) / 2; i >= 0; i-- { // 大顶堆
		downAdjust(arr, i, length-1)
	}
	// 进行堆排序
	for i := length - 1; i > 0; i-- {
		tmp := arr[i]
		arr[i] = arr[0]
		arr[0] = tmp
		downAdjust(arr, 0, i-1)
	}
}

func downAdjust(arr []int, parent, n int) {
	// 临时保存要下沉的元素; 定位左孩子节点的位置
	tmp, child := arr[parent], 2*parent+1
	for child <= n {
		if child+1 <= n && arr[child] < arr[child+1] { // 如果左孩子小于右孩子，则定位到右孩子
			child++
		}
		if arr[child] <= tmp { // 如果孩子节点小于或等于父节点，则下沉结束
			break
		}
		// 父节点进行下沉
		arr[parent] = arr[child]
		parent = child
		child = 2*parent + 1
	}
	arr[parent] = tmp
}
