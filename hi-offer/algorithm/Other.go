package algorithm

import (
	"math"
	"strconv"
	"strings"
)

/**
 * 002-替换空格
 *
 * 请实现一个函数，将一个字符串中的每个空格替换成“%20”。
 * 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
 *
 * 思路1:
 * 1. 从前往后遍历字符串，得到空格的数量。更新字符串的长度（扩容）。
 * 2. 从后往前替换字符串的话，每个字符串只需要移动一次；
 */
func ReplaceBlank20(src string) string {
	length := 0
	for _, v := range src {
		if ' ' == v {
			length += 3
			continue
		}
		length++
	}

	arr := make([]rune, length)
	i := 0
	for _, v := range src {
		arr[i] = v
		if ' ' == v {
			arr[i] = '%'
			arr[i+1] = '2'
			arr[i+2] = '0'
			i += 3
			continue
		}
		i++
	}
	return string(arr)
}

func ReplaceBlank201(src string) string {
	return strings.ReplaceAll(src, " ", "%20")
}

/**
 * 013-调整数组顺序使奇数位于偶数前面
 *
 * 输入一个整数数组，调整数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
 * 要求时间复杂度为O(n)，不要求保留数组元素的相对位置。
 *
 * 思路1:
 * 1. 两个指针，一个指针从前往后（找偶数），一个指针从后往前（找奇数），交换
 * 时间复杂度为O(n)
 */
func OddBeforeEven(arr []int) {
	if nil == arr {
		return
	}
	i, j := 0, len(arr)-1
	for i < j {
		for i < j && 0 != arr[i]%2 {
			i++
		}
		for i < j && 0 == arr[j]%2 {
			j--
		}
		if i < j {
			// swap
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp

			i++
			j--
		}
	}
}

func OddBeforeEven2(arr []int) {
	if nil == arr {
		return
	}
	length := len(arr)
	for i, j := -1, 0; j < length; j++ {
		if 0 != arr[j]%2 {
			i++
			// swap
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
		}
	}
}

/**
 * 只要不是特别大的内存开销，时间复杂度比较重要。因为改进时间复杂度对算法的要求更高
 *
 * 如果是顺序查找需要          O(N) 的时间；
 * 如果输入的是排序的数组则只需要 O(logN)的时间；
 * 如果事先已经构造好了哈希表，那查找在 O(1) 时间就能完成。
 *
 * 028-数组中出现次数超过一半的数字
 *
 * 思路1：
 * 题目给出的数组没有说是排序的，因此我们需要先给它排序。排序的时间复杂度是O(N*logN)。
 *
 * 思路2
 * 数组的特性：数组中有一个数字出现的次数超过了数组长度的一半。
 * 如果把这个数组『排序』，那么排序之后位于数组中间的数字一定就是那个出现次数超过数组长度一半的数字。
 * 也就是『中位数』
 * 有成熟的O(N)的算法得到数组中任意第k大的数字。
 *
 * 思路3
 * 前提：超过半数的数字存在
 * 数组中有一个数字出现的次数超过数组长度的一半，也就是说它出现的次数比其他所有数字出现次数的和还要多.
 * 因此我们可以考虑在遍历数组的时候保存两个值：一个是数组中的一个数字，一个是次数。
 * 1. 当我们遍历到下一个数字的时候，如果下一个数字和我们之前保存的数字相同，则次数加1；
 * 2. 如果下一个数字和我们之前保存的数字不同，则次数减1。
 * 3. 如果次数为零，我们需要保存下一个数字，并把次数设为1。
 *
 */
func FindMoreThanHalfNumber(arr []int) int {
	if nil == arr || len(arr) < 1 {
		return 0
	}
	length, may, cnt := len(arr), arr[0], 1
	for i := 1; i < length; i++ {
		if cnt < 0 {
			cnt = 1
			may = arr[i]
		}

		if may == arr[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return may
}

/**
 * 032-把数组排成最小的数
 *
 * 输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
 * 例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
 *
 * 思路1、全排列
 * 求出数组中所有数字的全排列，然后把每个全排列拼起来，求出拼出来的数字的最小值。
 *
 * 思路2、定义新的排序规则，排序后拼接
 * 例如：2和21 加起来比较221和212，212更小，于是说明21更小，把21放在2的前面
 */
func oddCompare(a, b int) bool {
	ab := strconv.Itoa(a) + strconv.Itoa(b)
	ba := strconv.Itoa(b) + strconv.Itoa(a)
	return ab >= ba
}

func oddPartition(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		//for left < right && arr[right] >= pivot {
		//	right--
		//}
		for left < right && oddCompare(arr[right], pivot) {
			right--
		}
		arr[left] = arr[right]
		//for left < right && arr[left] <= pivot {
		//	left++
		//}
		for left < right && oddCompare(pivot, arr[left]) {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

func OddQuickSort(arr []int, left, right int)  {
	if left < right {
		pivot := oddPartition(arr, left, right)
		OddQuickSort(arr, left, pivot-1)
		OddQuickSort(arr, pivot+1, right)
	}
}

func PrintOddQuickSortNumber(arr []int)  {
	if nil == arr {
		return
	}
	OddQuickSort(arr, 0, len(arr)-1)
	for _, v := range arr {
		print(v)
	}
	println()
}

/**
 * 033-丑数
 *
 * 把只包含质因子2、3和5的数称作丑数（Ugly Number）。
 * 例如6、8都是丑数，但14不是，因为它包含质因子7。
 * 习惯上我们把1当做是第一个丑数。
 * 求按从小到大的顺序的第N个丑数。
 *
 * 思路一：暴力搜索，时间复杂度不满足要求。
 * 思路二：动态规划方法，当前第n个丑数等于，前n-1个丑数中乘2，乘3，乘5中最小的数，且不再前n-1个丑数之中。
 */
/**
 * 思路二：不必维护三个队列
 * 维护三个指针
 * 分别作用为 乘2,3或者5，最开始大于原数组最大的指针。
 */
func GetUglyNumber(n int) int {
	if n < 1 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 1 // 习惯上我们把1当做是第一个丑数。
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x2, x3, x5 := arr[p2]*2, arr[p3]*3, arr[p5]*5
		uglyNow := multiIntMin(x2, x3, x5)
		arr[i] = uglyNow

		for x2 <= uglyNow {
			p2++
			x2 = arr[p2]*2
		}
		for x3 <= uglyNow {
			p3++
			x3 = arr[p3]*3
		}
		for x5 <= uglyNow {
			p5++
			x5 = arr[p5]*5
		}
	}
	return arr[n-1]
}

func multiIntMin(arr ...int) int {
	result := math.MaxInt64
	for _, v := range arr {
		if v < result {
			result = v
		}
	}
	return result
}
