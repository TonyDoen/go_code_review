package LCNOT

import (
	"container/list"
	"math"
	"sort"
)

/**
 * 区间相关问题
 *
 * 区间问题就是线段问题，让你合并所有线段，找出线段的交集，
 * 主要2个技巧
 * <1>排序;
 * <2>画图;
 *
 *
 */
/**
 * 区间覆盖问题
 * LeetCode 1288 删除被覆盖区间
 *
 * 给你一个区间列表，请你删除列表种被其他区间所覆盖的区间。
 * 只有当 c <= a 且 b <= d 时，我们才认为区间 [a, b) 被左闭右开区间 [c, d) 覆盖
 * 在完成所有删除操作后，请你返回列表中剩余区间的数目
 *
 *
 * 例子1：
 * input: intervals = [[1,4],[3,6],[2,8]]
 * output: 2
 * explain: 区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除。
 *
 *
 * 思路：
 * 去掉被覆盖区间后，还剩下多少区间，总区间数-被覆盖区间数=剩余区间数
 * 排序区间后
 * 有3种相对位置
 * <1>
 *     |______________|
 *         |_______|
 *
 * <2>
 *     |______________|
 *               |___________|
 *
 * <3>
 *     |___________|
 *                     |____________|
 *
 * 情况1， 找到覆盖区域。
 * 情况2，2个区域可以合并，成为一个大区域
 * 情况3，2个区域完全不相交
 *
 */
func RemoveCoveredIntervals(intvs [][]int) int {
	// 按照起点升序排列，起点相同同时按照终点降序排列
	// 保证长的区域在上面，这样才会被判定为覆盖。
	// 复用上面的信封封装逻辑
	var es, err = NewEnvelopes(intvs)
	if nil != err {
		return -1
	}
	sort.Sort(es)
	intvs = es.EnvelopeHW

	// 记录合并区域的起点和终点
	var left, right, rs, length = intvs[0][0], intvs[0][1], 0, len(intvs)
	for i := 1; i < length; i++ {
		var intv = intvs[i]
		// 情况1， 找到覆盖区域。
		if left <= intv[0] && right >= intv[1] {
			rs++
		}
		// 情况2， 2个区域可以合并，成为一个大区域
		if right >= intv[0] && right <= intv[1] {
			right = intv[1]
		}
		// 情况3， 2个区域完全不相交
		if right < intv[0] {
			left = intv[0]
			right = intv[1]
		}
	}
	return length - rs
}

/**
 * 区间合并问题
 * LeetCode 56
 *
 * 给出一个区间的集合，请合并所有重叠的区间
 * 例子1：
 * input:  [[1, 3], [2, 6], [8, 10], [15, 18]]
 * output: [[1, 6], [8, 10], [15, 18]]
 * explain:[1, 3] 和 [2, 6] 重叠，将它们合并 [1, 6]
 *
 * 例子2：
 * input:  [[1, 4], [4, 5]]
 * output: [[1, 5]]
 * explain:[1, 4] 和 [4, 5] 重叠，将它们合并 [1, 5]
 *
 */
func MergeCoveredIntervals(intvs [][]int) *list.List {
	// 按照起点升序排列，起点相同同时按照终点降序排列
	// 保证长的区域在上面，这样才会被判定为覆盖。
	var es, err = NewEnvelopes(intvs)
	if nil != err {
		return nil
	}
	sort.Sort(es)
	intvs = es.EnvelopeHW

	var rs, length = list.New(), len(intvs)
	rs.PushBack(intvs[0])

	for i := 1; i < length; i++ {
		var curr = intvs[i]
		// rs's last element link
		var last = rs.Back().Value.([]int)
		if curr[0] <= last[1] {
			// find the max end
			last[1] = int(math.Max(float64(last[1]), float64(curr[1])))
		} else {
			// handle the next area
			rs.PushBack(curr)
		}
	}
	return rs
}

/**
 * 区间交集问题
 * LeetCode 986
 *
 * 给定2个由一些闭区间组成的列表，每个区间列表都是成对不相交，并且已经排序。
 * 返回这2个区间列表的交集。
 *
 * 形式上，闭区间 [a, b] 其中 a <= b, 表示 实数x的集合，而 a <= x <= b。2个闭区间的交集是一组实数，要么为空集，要么为闭区间。
 * 例如，[1, 3] 和 [2, 4] 的交集为 [2, 3]
 *
 * 例子1
 * A：  |_________|     |___________|        |_______________________|    |_____|
 * B：        |_________|       |________|                      |_________|     |_____________|
 * ans:       |___|     |       |___|                           |____|    |     |
 *
 * input: A = [[0, 2], [5, 10], [13, 23], [24, 25]]
 *        B = [[1, 5], [8, 12], [15, 24], [25, 26]]
 * output:[[1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]]
 *
 * 注意：输入和所需要的输出都是区间对象组成的列表，而不是数组或者列表
 *
 * 1. 不相交情况
 * <1>  |_________|                                     <2>                  |_________|
 *     a0        a1                                                          a0        a1
 *                    |_________|                           |_________|
 *                    b0        b1                          b0        b1
 *
 * 2. 取交集的情况
 *    4种情况
 *
 * 3. i, j 更新步骤
 *    |_________|     |___________|
 *     i
 *
 *          |_________|       |________|
 *           j
 */
func IntervalIntersection(a, b [][]int) *list.List {
	var rs = list.New()
	for i, j, aLength, bLength := 0, 0, len(a), len(b); i < aLength && j < bLength; {
		var a0, a1, b0, b1 = a[i][0], a[i][1], b[j][0], b[j][1]
		if b1 < a0 || a1 < b0 {
			// nothing
		} else {
			// Intersection
			rs.PushBack([]int{int(math.Max(float64(a0), float64(b0))), int(math.Max(float64(a1), float64(b1)))})
		}

		// i, j update
		if b1 < a1 {
			j++
		} else {
			i++
		}
	}
	return rs
}
