package algorithm

/**
 * 001-二维数组查找
 *
 * 题目：
 * 二维数组中的查找
 *
 * 题目描述：
 * 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
 * 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。（每个一维数组的长度相同）
 * eg:
 * 1 3 5 7
 * 2 4 6 8
 * 3 5 7 9
 *
 * 思路：
 * 1、根据数组特点，从右上角位置开始查找(右上角元素是所在行的最大值，所在列的最小值。)
 * 2、如果相等，返回true即可
 * 3、如果目标值大于当前值，忽略本行，继续查找
 * 4、如果目标值小于当前值，忽略本列，继续查找
 *
 * 时间复杂度：O（row+col）
 */
func TwoDimensionalArrayLookup(arr [][]int, target int) bool {
	if nil == arr || len(arr) < 1 || len(arr[0]) < 1 {
		return false
	}
	row, col := len(arr), len(arr[0])
	for i, j := 0, col-1; i < row && j >= 0; {
		curVal := arr[i][j]
		if target == curVal {
			return true
		} else if curVal < target {
			i++
		} else {
			j--
		}
	}
	return false
}

/**
 * 006-旋转数组的最小数字（二分查找）
 *
 * 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
 * 输入一个非减排序的数组（即，递增的数组）的一个旋转，输出旋转数组的最小元素。
 * 例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
 *
 * NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
 *
 * 时间复杂度：O(logN)
 */
func MoveLeftFindMin(arr []int) int {
	if nil == arr || len(arr) < 1 {
		return 0
	}
	left, right := 0, len(arr)-1
	if arr[left] < arr[right] {
		return arr[0]
	}
	mid := 0
	for arr[left] >= arr[right] {
		if 1 == right-left {
			mid = right
			break
		}
		mid = left + (right-left)/2
		//下标为left,right,mid指向的元素相等，无法明确待缩减的范围，采用顺序查找。
		if arr[mid] == arr[left] && arr[mid] == arr[right] {
			min := arr[left]
			for i := left + 1; i < right; i++ {
				if min > arr[i] {
					min = arr[i]
				}
			}
			return min
		}

		if arr[mid] >= arr[left] {
			left = mid
		} else if arr[mid] <= arr[right] {
			right = mid
		}
	}
	return arr[mid]
}

/**
 * 037-数字在排序数组中出现的次数（二分查找）。统计一个数字在排序数组中出现的次数。
 *
 * 思路：
 * 在有序数组中的查找，二分查找具有时间复杂度上的优势。
 * 这里直接利用二分查找来分别查找目标值的左右边界来确定出现的次数，那么时间复杂度为O(logN)。
 */
func GetTargetCnt(arr []int, target int) int {
	if nil == arr || len(arr) < 1 {
		return 0
	}
	first := firstTarget(arr, target)
	last := lastTarget(arr, target)
	if -1 != first && -1 != last {
		return last - first + 1
	}
	return 0
}
func firstTarget(arr []int, target int) int { // arr: 递增
	length := len(arr)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target || (mid-1 >= 0 && arr[mid-1] == target) {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func lastTarget(arr []int, target int) int { // arr: 递增
	length := len(arr)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target || (mid+1 < length && arr[mid+1] == target) {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
