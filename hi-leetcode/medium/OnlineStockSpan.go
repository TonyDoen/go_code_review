package medium

import (
	"container/list"
)

/**
 * Write a class StockSpanner which collects daily price quotes for some stock, and returns the span of that stock's price for the current day.
 * The span of the stock's price today is defined as the maximum number of consecutive days (starting from today and going backwards) for which the price of the stock was less than or equal to today's price.
 * For example, if the price of a stock over the next 7 days were [100, 80, 60, 70, 60, 75, 85], then the stock spans would be [1, 1, 1, 2, 1, 4, 6].
 *
 * Example 1:
 * Input: ["StockSpanner","next","next","next","next","next","next","next"], [[],[100],[80],[60],[70],[60],[75],[85]]
 * Output: [null,1,1,1,2,1,4,6]
 * Explanation:
 * First, S = StockSpanner() is initialized.  Then:
 * S.next(100) is called and returns 1,
 * S.next(80) is called and returns 1,
 * S.next(60) is called and returns 1,
 * S.next(70) is called and returns 2,
 * S.next(60) is called and returns 1,
 * S.next(75) is called and returns 4,
 * S.next(85) is called and returns 6.
 * Note that (for example) S.next(75) returned 4, because the last 4 prices
 * (including today's price of 75) were less than or equal to today's price.
 *
 * Note:
 * Calls to StockSpanner.next(int price)will have 1 <= price <= 10^5.
 * There will be at most 10000 calls to StockSpanner.next per test case.
 * There will be at most 150000 calls to StockSpanner.next across all test cases.
 * The total time limit for this problem has been reduced by 75% for C++, and 50% for all other languages.
 */

/**
 * 题意：股票价格跨度
 * 这道题定义了一个 StockSpanner 的类，有一个 next 函数，每次给当天的股价，让我们返回之前连续多少天都是小于等于当前股价。
 *
 * 跟 OJ 抗争多年的经验告诉我们，不要想着可以用最暴力的向前搜索的方法，这种解法太 trivial 了，肯定无法通过的。那么可以找连续递增的子数组的长度么，其实也是不行的，就拿题目中的例子来说吧 [100, 80, 60, 70, 60, 75, 85]，数字 75 前面有三天是比它小的，但是这三天不是有序的，是先增后减的，那怎么办呢？
 *
 * 思路：
 * 我们先从简单的情况分析，假如当前的股价要小于前一天的，那么连续性直接被打破了，所以此时直接返回1就行了。
 * 但是假如大于等于前一天股价的话，情况就比较 tricky 了，因为此时所有小于等于前一天股价的天数肯定也是小于等于当前的，那么我们就需要知道是哪一天的股价刚好大于前一天的股价，然后用这一天的股价跟当前的股价进行比较，若大于当前股价，说明当前的连续天数就是前一天的连续天数加1，而若小于当前股价，我们又要重复这个过程，去比较刚好大于之前那个的股价。所以我们需要知道对于每一天，往前推刚好大于当前股价的是哪一天，用一个数组 pre，其中 pre[i] 表示从第i天往前推刚好大于第i天的股价的是第 pre[i] 天。
 *
 * 接下来看如何实现 next 函数，首先将当前股价加入 nums 数组，然后前一天在数组中的位置就是 (int)nums.size()-2。再来想想 corner case 的情况，假如当前是数组中的第0天，前面没有任何股价了，我们的 pre[0] 就赋值为 -1 就行了，怎么知道当前是否是第0天，就看 pre 数组是否为空。再有就是由于i要不断去 pre 数组中找到之前的天数，所以最终i是有可能到达 pre[0] 的，那么就要判断当i为 -1 时，也要停止循环。循环的最后一个条件就是当之前的股价小于等当前的估计 price 时，才进行循环，这个前面讲过了，循环内部就是将 pre[i] 赋值给i，这样就完成了跳到之前天的操作。while 循环结束后要将i加入到 pre 数组，因为这个i就是从当前天往前推，一个大于当前股价的那一天，有了这个i，就可以计算出连续天数了，参见代码如下：
 */

type StockSpanner struct {
	num []int // 切片, slice
	pre []int
}

func (ss *StockSpanner) next(price int) int {
	if nil == ss {
		return -1
	}

	ss.num = append(ss.num, price)
	i := len(ss.num) - 2
	for len(ss.pre) > 0 && i >= 0 && ss.num[i] <= price {
		i = ss.pre[i]
	}
	ss.pre = append(ss.pre, i)
	return len(ss.pre) - 1 - i
}

type Pair struct {
	k interface{}
	v interface{}
}

func NewPair(k, v interface{}) *Pair {
	return &Pair{k, v}
}

type StockSpanner2 struct {
	st list.List
}

func (ss *StockSpanner2) next(price int) int {
	if nil == ss {
		return -1
	}
	cnt := 1
	for ss.st.Len() > 0 {
		front := ss.st.Front()
		//if nil == front.Value {
		//	break
		//}
		pair := front.Value.(*Pair)
		//if nil == pair {
		//	break
		//}
		if pair.k.(int) <= price {
			cnt += pair.v.(int)
			ss.st.Remove(front)
		} else {
			break
		}
	}
	ss.st.PushFront(NewPair(price, cnt)) // stack
	return cnt
}
