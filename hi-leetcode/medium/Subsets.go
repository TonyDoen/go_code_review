package medium

import (
	"container/list"
)

/**
 * url: https://www.cnblogs.com/grandyang/p/4309345.html
 * Given a set of distinct integers, S, return all possible subsets.
 *
 * Note:
 * Elements in a subset must be in non-descending order.
 * The solution set must not contain duplicate subsets.
 *
 * For example,
 * If S = [1,2,3], a solution is:
 * [
 *   [3],
 *   [1],
 *   [2],
 *   [1,2,3],
 *   [1,3],
 *   [2,3],
 *   [1,2],
 *   []
 * ]
 *
 *
 * 题意：这道求子集合的问题,由于其要列出所有结果,
 * 思路：按照以往的经验,肯定要是要用递归来做。
 * 这道题其实它的非递归解法相对来说更简单一点,下面我们先来看非递归的解法。
 *
 * 我们可以一位一位的往上叠加,比如对于题目中给的例子 [1,2,3] 来说,
 * 1. 最开始是空集,                                         []
 * 2. 那么我们现在要处理1,就在空集上加1,为 [1],现在我们有两个      [],[1],
 * 3. 下面我们来处理2,我们在之前的子集基础上,每个都加个2,可以得到    [], [1], [2], [1, 2],
 * 4. 同理处理3的情况可得                                     [], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3],
 */
func GetSubsets00(arr []int) *list.List {
	result := list.New()
	result.PushBack(list.New())
	QuickSort(arr)
	for _, v := range arr {
		cur := result.Front()
		for nil != cur {
			clone := cloneList(cur.Value.(*list.List))
			clone.PushBack(v)
			result.PushFront(clone) // NOTICE
			cur = cur.Next()
		}
	}
	return result
}

func cloneList(src *list.List) *list.List {
	cur := src.Front()
	result := list.New()
	for nil != cur {
		result.PushBack(cur.Value)
		cur = cur.Next()
	}
	return result
}

/**
 * url: https://www.cnblogs.com/grandyang/p/4310964.html
 * Given a collection of integers that might contain duplicates, S, return all possible subsets.
 *
 * Note:
 * Elements in a subset must be in non-descending order.
 * The solution set must not contain duplicate subsets.
 *
 * For example,
 * If S = [1,2,2], a solution is:
 * [
 *   [2],
 *   [1],
 *   [1,2,2],
 *   [2,2],
 *   [1,2],
 *   []
 * ]
 *
 *
 * 题意：这道求子集合的问题,由于其要列出所有结果,
 * 思路：按照以往的经验,肯定要是要用递归来做。
 * 这道题其实它的非递归解法相对来说更简单一点,下面我们先来看非递归的解法。
 *
 * 我们可以一位一位的往上叠加,比如对于题目中给的例子 [1,2,3] 来说,
 * 1. 最开始是空集,                                         []
 * 2. 那么我们现在要处理1,就在空集上加1,为 [1],现在我们有两个      [],[1],
 * 3. 下面我们来处理2,我们在之前的子集基础上,每个都加个2,可以得到    [], [1], [2], [1, 2],
 * 4. 同理处理3的情况可得                                     [], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3],
 */
func GetSubsets10(arr []int) *list.List {
	result := list.New()
	result.PushBack(list.New())
	QuickSort(arr)
	for _, v := range arr {
		cur := result.Front()
		for nil != cur {
			// clone
			clone := cloneList(cur.Value.(*list.List))
			clone.PushBack(v)
			// compare duplicate
			if containList(result, clone) {
				cur = cur.Next()
				continue
			}
			result.PushFront(clone) // NOTICE
			cur = cur.Next()
		}
	}
	return result
}

func containList(outer, inner *list.List) bool {
	outCur := outer.Front()
	for nil != outCur {
		if compareList(outCur.Value.(*list.List), inner) {
			return true
		}
		outCur = outCur.Next()
	}
	return false
}

func compareList(a, b *list.List) bool { // true
	if nil == a && nil == b {
		return true
	}
	if nil == a || nil == b {
		return false
	}
	aCur, bCur := a.Front(), b.Front()
	for nil != aCur && nil != bCur {
		if aCur.Value.(int) != bCur.Value.(int) {
			return false
		}
		aCur = aCur.Next()
		bCur = bCur.Next()
	}
	return aCur == bCur
}
