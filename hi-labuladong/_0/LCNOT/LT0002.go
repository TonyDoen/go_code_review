package LCNOT

import "sort"

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

