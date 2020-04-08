package medium

import (
	"container/list"
	"strconv"
)

/**
 * url: https://www.cnblogs.com/grandyang/p/4682458.html
 *
 * Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators.
 * The valid operators are +, - and *.
 *
 * Example 1:
 * Input: "2-1-1"
 * Output: [0, 2]
 *
 * Explanation:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 * Example 2:
 * Input: "2*3-4*5"
 * Output: [-34, -14, -10, -10, 10]
 *
 * Explanation:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 * 题意：添加括号的不同方式，求不同括号添加方式的计算结果
 * 思路：若 input 是数字和运算符的时候，比如 "1+1" 这种情况，那么加不加括号也没有任何影响，因为只有一个计算，结果一定是2。
 * 复杂一点的话，input 是 "2-1-1" 时，就有两种情况了，(2-1)-1 和 2-(1-1)，由于括号的不同，得到的结果也不同，
 * 但如果我们把括号里的东西当作一个黑箱的话，那么其就变为 ()-1  和 2-()，其最终的结果跟括号内可能得到的值是息息相关的，
 * 那么再概括一点，实际上就可以变成 () ? () 这种形式，两个括号内分别是各自的表达式，最终会分别计算得到两个整型数组，中间的问号表示运算符，可以是加，减，或乘。
 *
 * 那么问题就变成了从两个数组中任意选两个数字进行运算，
 * 这种左右两个括号代表的黑盒子就交给递归去计算，
 * 像这种分成左右两坨的 pattern 就是分治法 Divide and Conquer 了。
 *
 */
func DifferentWaysToCompute(input string) *list.List {
	length, result := len(input), list.New()
	if length < 1 {
		return result
	}
	for i := 0; i < length; i++ {
		ci := input[i]
		if '+' == ci || '-' == ci || '*' == ci {
			left := DifferentWaysToCompute(input[0:i])
			right := DifferentWaysToCompute(input[i+1 : length])
			for lCur := left.Front(); nil != lCur; lCur = lCur.Next() {
				for rCur := right.Front(); nil != rCur; rCur = rCur.Next() {
					lVal, rVal := lCur.Value.(int), rCur.Value.(int)
					if '+' == ci {
						result.PushBack(lVal + rVal)
					} else if '-' == ci {
						result.PushBack(lVal - rVal)
					} else if '*' == ci {
						result.PushBack(lVal * rVal)
					}
				}
			}
		}
	}
	if result.Len() < 1 {
		v, _ := strconv.Atoi(input)
		result.PushBack(v)
	}
	return result
}

// 函数参数: 裸函数作为参数
type DifferentWaysToComputeTemplate func(input string) *list.List

func doDifferentWaysToCompute(result *list.List, input string, template DifferentWaysToComputeTemplate) {
	length := len(input)
	if length < 1 {
		return
	}
	for i := 0; i < length; i++ {
		ci := input[i]
		if '+' == ci || '-' == ci || '*' == ci {
			left := template(input[0:i])
			right := template(input[i+1 : length])
			for lCur := left.Front(); nil != lCur; lCur = lCur.Next() {
				for rCur := right.Front(); nil != rCur; rCur = rCur.Next() {
					lVal, rVal := lCur.Value.(int), rCur.Value.(int)
					if '+' == ci {
						result.PushBack(lVal + rVal)
					} else if '-' == ci {
						result.PushBack(lVal - rVal)
					} else if '*' == ci {
						result.PushBack(lVal * rVal)
					}
				}
			}
		}
	}
}

func DifferentWaysToCompute1(input string) *list.List {
	result := list.New()
	doDifferentWaysToCompute(result, input, DifferentWaysToCompute1)
	if result.Len() < 1 {
		v, _ := strconv.Atoi(input)
		result.PushBack(v)
	}
	return result
}

var cacheMap = make(map[string]*list.List)

func DifferentWaysToCompute2(input string) *list.List {
	// 1.1 优化 避免递归中的重复计算
	tmp := cacheMap[input]
	if nil != tmp {
		return tmp
	}

	result := list.New()
	doDifferentWaysToCompute(result, input, DifferentWaysToCompute2)
	if result.Len() < 1 {
		v, _ := strconv.Atoi(input)
		result.PushBack(v)
	}
	// 1.2 优化 避免递归中的重复计算
	cacheMap[input] = result
	return result
}
