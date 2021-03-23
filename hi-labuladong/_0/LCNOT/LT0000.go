package LCNOT

/**
 * 此算法有什么缺陷?
 * 比如说给你有序数组 nums = [1,2,2,2,3] , target 为 2,
 *
 * 希望返回的左侧边界,即索引 1,
 * 或者希望右侧边界,即索引 3,
 * 这样的话此算法是无法处理的。
 */
func BinarySearch0(nums []int, target int) int {
	var left, right = 0, len(nums) - 1 // careful
	for ; left <= right; {             // careful
		var mid = left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1 // careful
		} else if target < nums[mid] {
			right = mid - 1 // careful
		}
	}
	return -1
}

/**
 *
 * 寻找左侧边界的二分搜索
 *
 * 为什么该算法能够搜索左侧边界?
 * 关键在于对于 nums[mid] == target 这种情况的处理
 * if (target == nums[mid]) {
 *     right = mid;
 * }
 * 找到 target 时不要立即返回,而是缩小「搜索区间」的上界 right. 在区间 [left, mid) 中继续搜索,即不断向左收缩,达到锁定左侧边界的目的
 *
 */
func LeftBinarySearch0(nums []int, target int) int {
	var length = len(nums)
	if 0 == length {
		return -1
	}
	var left, right = 0, length
	for ; left < right; { // careful
		var mid = left + (right-left)/2
		if target == nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid // careful
		}
	}

	if left >= length || target != nums[left] {
		return -1
	}
	return left
}

/**
 * 寻找右侧边界的二分查找
 *
 * 为什么这个算法能够找到右侧边界?
 * if (target == nums[mid]) {
 *     left = mid + 1;
 * }
 * 当 nums[mid] == target 时,不要立即返回,而是增大「搜索区间」的下界 left, 使得区间不断向右收缩,达到锁定右侧边界的目的
 */
func RightBinarySearch0(nums []int, target int) int {
	var length = len(nums)
	if 0 == length {
		return -1
	}

	var left, right = 0, length // careful
	for ; left < right; {
		var mid = left + (right-left)/2
		if target == nums[mid] {
			left = mid + 1 // careful
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		}
	}

	return left - 1
}

/**
 * 统一模板二分查找法
 */
func NormalBinarySearch1(nums []int, target int) int {
	var length = len(nums)
	var left, right = 0, length - 1
	for ; left <= right; {
		var mid = left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			return mid
		}
	}
	return -1
}

func LeftBinarySearch1(nums []int, target int) int {
	var length = len(nums)
	var left, right = 0, length - 1
	for ; left <= right; {
		var mid = left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			right = mid - 1 // not return, locate left bound
		}
	}

	// check left bound
	if left >= length || target != nums[left] {
		return -1
	}
	return left
}

func RightBinarySearch1(nums []int, target int) int {
	var length = len(nums)
	var left, right = 0, length - 1
	for ; left <= right; {
		var mid = left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			left = mid + 1 // not return, locate right bound
		}
	}

	// check right bound
	if right < 0 || target != nums[right] {
		return -1
	}
	return right
}
