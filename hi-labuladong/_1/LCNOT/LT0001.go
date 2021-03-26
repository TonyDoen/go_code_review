package LCNOT

import (
	"errors"
	"sort"
)

/**
 * 很多算法问题都需要排序技巧，其难点不在于排序本身，而是需要巧妙地排序进行预处理，将算法问题进行转换，为之后的操作打下基础。
 * 信封嵌套问题就需要先按特定的规则排序，之后就转换为一个最长递增子序列问题，可以用前文 动态规划设计之最长递增子序列 的技巧来解决了。
 *
 * 信封嵌套问题是个很有意思且经常出现在生活中的问题
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式(w, h)出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就能放到另一个信封里，如同俄罗斯套娃。
 * 请计算最多能有多少个信封组成俄罗斯套娃信封。
 * 说明：
 * 不允许旋转信封
 *
 * 例子1：
 * input:  envelopes = [[5, 4], [6, 4], [6, 7], [2, 3]]
 * output: 3
 * explain:最多信封个数3， 组合：[2, 3] => [5, 4] => [6, 7]
 *
 *
 * 思路：
 * 这是 最长递增子序列(Longest Increasing Subsequence, LIS)的一个变种问题。
 * 每次合法的嵌套是大的套小的，相当于找一个最长递增的子序列，其长度就是最多能嵌套的信封个数。
 *
 * 难点在与，标准的 LIS 算法只能在数组中寻找最长子序列，而我们的信封是由(w, h) 这样的二维数对表示，如何迁移 LIS 算法到二维数对？
 * <1> 通过 w [x] h 计算面积是不行的。 1x10 > 3x3 但是这2个信封无法相互嵌套。
 * <2> 巧妙 的先对宽度 w 进行升序排序，如果遇到 w 相同的情况，按照高度 h 降序排列。之后把所有 h 作为一个数组，在这个数组里计算 LIS 的长度，即答案。
 *
 * 解释 LIS：
 * 令 arr = [5, 1, 4, 3, 4, 2];
 *              *     *  *
 * 则 LIS = [1, 3, 4]
 * len(LIS) = 3
 *
 * 解释 解法2：
 *  | 宽度 w,     高度 h  |
 *  | [ 2,         3 ]   |
 *  | [ 5,         4 ]   |
 *  | [ 6,         7 ]   |
 *  | [ 6,         4 ]   |
 *  |                    |
 * \/                   \/
 * 升序                  降序
 *
 * 这个问题是个 Hard 级别的题目，难就难在排序，正确地排序后此问题就被转化成了一个标准的 LIS 问题，容易解决一些。
 * 为了清晰，我将代码分为了两个函数， 你也可以合并，这样可以节省下height数组的空间。
 *
 * 此算法的时间复杂度为O(NlogN)，因为排序和计算 LIS 各需要 O(NlogN)的时间。
 * 空间复杂度为 O(N), 因为计算 LIS 的函数需要一个 top 数组。
 *
 * 衍生问题
 * 这个问题还可以拓展到3维，比如把嵌套信封变成嵌套箱子，计算最多能嵌套多少箱子？
 * <1> 如果按照信封题目的思路，固定[长, 宽]; 在第三维度[高] 上找 LIS， 这个思路是有问题的，问题是怎么固定[长, 宽]。
 * <2> 实际上这类问题叫做[偏序问题]，上升到3维，难度一下就提升了。这需要借助 [树状数组]
 *
 */
func MaxEnvelopes(envelopes [][]int) int {
	var es, err = NewEnvelopes(envelopes)
	if nil != err {
		return -1
	}
	sort.Sort(es)

	var sortEs = es.EnvelopeHW
	var length = len(sortEs)
	var h = make([]int, length)
	for i := 0; i < length; i++ {
		h[i] = sortEs[i][1]
	}
	return lengthOfLIS(h)
}

// implement Length of LIS
func lengthOfLIS(nums []int) int {
	var piles, n = 0, len(nums)
	var top = make([]int, n)
	for i := 0; i < n; i++ {
		var p = nums[i] // this digit is needed to handle
		var left, right = 0, piles
		// binary search insertion position
		for left < right {
			var mid = (left + right) / 2
			if top[mid] >= p {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == piles {
			piles++
		}
		// put this digit on the top of heap
		top[left] = p
	}
	// size of heap is the LIS's length
	return piles
}

// implement of sort's Interface
type Envelopes struct {
	/**
	 * envelopeHW => {{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	 * index      =>   0,      1,      2,      3,
	 *
	 * explain    => 0's envelope. width is 5, height is 4
	 *
	 */
	EnvelopeHW [][]int
}

func (es *Envelopes) Len() int {
	//if nil == es || nil == es.envelopeHW{
	//	return -1
	//}
	return len(es.EnvelopeHW)
}

func (es *Envelopes) Less(i, j int) bool {
	//if nil == es || nil == es.envelopeHW {
	//	return false
	//}
	// 按照宽度升序。若宽度相同，则按照高度降序
	var rs = 0
	var iv = es.EnvelopeHW[i]
	var jv = es.EnvelopeHW[j]
	if iv[0] == jv[0] {
		rs = jv[1] - iv[1]
	} else {
		rs = iv[0] - jv[0]
	}

	return rs < 0
}

func (es *Envelopes) Swap(i, j int) {
	var iv = es.EnvelopeHW[i]
	es.EnvelopeHW[i] = es.EnvelopeHW[j]
	es.EnvelopeHW[j] = iv
}

func NewEnvelopes(envelopes [][]int) (*Envelopes, error) {
	if nil == envelopes {
		return nil, errors.New("param is error")
	}
	return &Envelopes{EnvelopeHW: envelopes}, nil
}
