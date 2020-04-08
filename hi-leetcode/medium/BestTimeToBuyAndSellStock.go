package medium

import "math"

/**
 * Say you have an array for which the i th element is the price of a given stock on day i.
 * If you were only permitted to complete at most one transaction
 * (ie, buy one and sell one share of the stock),
 * design an algorithm to find the maximum profit.
 *
 * 题意：
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票），
 * 设计一个算法来计算你所能获取的最大利润。
 * 注意你不能在买入股票前卖出股票。
 *
 * 示例 1:
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 *      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 *
 * 示例 2:
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 * 思路1：(贪心算法)
 * 1、初始化数组的第一个元素为最低价格
 * 2、从左到右遍历，更新最低价格，更新最大收益值
 *
 */
func GetBestTimeToBuyAndSellStock00(prices []int) int {
	if nil == prices || len(prices) < 2 {
		return 0
	}
	min, maxProfit := prices[0], 0
	for _, cur := range prices {
		if cur < min {
			min = cur
		} else {
			curProfit := cur - min
			if curProfit > maxProfit {
				maxProfit = curProfit
			}
		}
	}
	return maxProfit
}

/**
 * 思路2：
 * 1、用sell表示初始时的利润为0，buy表示最便宜股票的价格
 * 2、从左到右遍历，buy表示前些天买入最便宜股票的股价
 * 3、sell保存前些天买入最便宜股票后再在股票最高时卖出股票的最大利润
 * 延伸：
 * 此思路可类推进行两次交易的最大利益。
 */
func GetBestTimeToBuyAndSellStock01(prices []int) int {
	if nil == prices || len(prices) < 2 {
		return 0
	}
	buy, maxProfit := -prices[0], 0
	for _, cur := range prices {
		if -cur > buy { // 更新最低价格
			buy = -cur
		}
		curProfit := cur + buy
		if curProfit > maxProfit { // 更新利润
			maxProfit = curProfit
		}
	}
	return maxProfit
}

/**
 * Say you have an array for which the i th element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit.
 * You may complete as many transactions as you like
 * (ie, buy one and sell one share of the stock multiple times).
 *
 * However, you may not engage in multiple transactions at the same time
 * (ie, you must sell the stock before you buy again).
 *
 * 题意：
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 * 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 * 示例 1:
 * 输入: [7,1,5,3,6,4]
 * 输出: 7
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，
 * 在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，
 * 在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 *
 * 示例 2:
 * 输入: [1,2,3,4,5]
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，
 * 在第 5 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 * 示例 3:
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 * 思路：
 * 1、初始化收益为0
 * 2、从左到右遍历，发现当前元素比刚遍历的元素大，即可取得收益，累加结果即为最大收益
 */
func GetBestTimeToBuyAndSellStock10(prices []int) int {
	if nil == prices || len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	for i, length := 1, len(prices); i < length; i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

/**
 * Say you have an array for which the i th element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit.
 * You may complete at most two transactions.
 *
 * Note:
 * You may not engage in multiple transactions at the same time
 * (ie, you must sell the stock before you buy again).
 *
 * 题意：（与上面的区别是只能完成 2 笔交易）
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 2 笔交易。
 *
 * 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 * 思路：(贪心算法)
 * 1、用 sell1 表示初始时的利润为0，
 * buy1 表示最便宜股票的价格，
 * 用 sell2 表示交易两次的利润，
 * buy2 表示第一次售出股票后，再买入后面某一天股票后的收益
 * 2、从左到右遍历，buy1表示前些天买入最便宜股票的股价
 * sell1保存前些天买入最便宜股票后再在股票最高时卖出股票的最大利润
 * 3、buy2表示第一次售出股票后，再买入后面某一天股票后的净收益
 * sell2表示二次买卖或者一次买卖的最大收益
 * (buy2之前的净收益+curPrice今天卖出股票后收益)
 */
func GetBestTimeToBuyAndSellStock20(prices []int) int {
	if nil == prices || len(prices) < 2 {
		return 0
	}
	buy1, sell1, buy2, sell2 := math.MinInt64, 0, math.MinInt64, 0
	for _, cur := range prices {
		if -cur > buy1 { // 最便宜的股票价格
			buy1 = -cur
		}
		curBuy1 := buy1 + cur // 一次交易的最大收益
		if curBuy1 > sell1 {
			sell1 = curBuy1
		}

		curSell2 := sell1 + (-cur) // 之前天先进行第一次交易后，在买入今天股票后的净利润
		if curSell2 > buy2 {
			buy2 = curSell2
		}
		curBuy2 := buy2 + cur // 二次交易的收益(卖出今天股票后的收益)
		if curBuy2 > sell2 {
			sell2 = curBuy2
		}
	}
	return sell2
}

/**
 * Say you have an array for which the i th element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit.
 * You may complete at most k transactions.
 * (ie, buy one and sell one share of the stock k times).
 *
 * However, you may not engage in multiple transactions at the same time
 * (ie, you must sell the stock before you buy again).
 *
 * 题意：（与上面的区别是只能完成 k 笔交易）
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 * 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 * 示例 1:
 * 输入: [7,1,5,3,6,4]; k=2
 * 输出: 7
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，
 * 在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，
 * 在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 *
 * 示例 2:
 * 输入: [1,2,3,4,5]; k=2
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，
 * 在第 5 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 * 示例 3:
 * 输入: [7,6,4,3,1]; k=2
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 * 思路：动态规划
 *
 * 状态穷举：
 * dp[i][k][1] 标识第i天 交易了k次手里持有股票的收益 ；
 * dp[i][k][0] 标识第i天 交易了k次手里没有股票的收益
 *
 * 状态转移：(买入股票的时候k+1)
 * dp[i][k][1]怎么获得，第i-1天持有股票没有卖出 dp[i-1][k][1] || 第i-1天没有股票买入dp[i-1][k-1][0]-prices[i]
 * dp[i][k][0]怎么获得，第i-1天没有股票没有卖入 dp[i-1][k][0] || 第i-1天有股票卖出 dp[i-1][k][0]+prices[i]
 *
 * base case：
 * dp[-1][k][0] = dp[i][0][0] = 0 (没开始 或者 没有交易，手里没有股票)
 * dp[-1][k][1] = dp[i][0][1] = -infinity (没开始，或没交易， 手里有股票，因为此种情况不可能)
 *
 * 优化后解法如下：
 * t[i][0] 和 t[i][1]分别表示第i比交易买入和卖出时 各自的最大收益
 *
 * 注意：
 * 当k大于等于数组长度一半时, 问题退化为贪心问题此时采用 买卖股票的最佳时机 II
 * 的贪心方法解决可以大幅提升时间性能, 对于其他的k, 可以采用 买卖股票的最佳时机 III
 * 的方法来解决, 在III中定义了两次买入和卖出时最大收益的变量, 在这里就是k租这样的
 * 变量, 即问题IV是对问题III的推广, t[i][0]和t[i][1]分别表示第i比交易买入和卖出时
 * 各自的最大收益
 */
func GetBestTimeToBuyAndSellStock30(prices []int, k int) int {
	if nil == prices || len(prices) < 2 {
		return 0
	}
	length := len(prices)
	if k > length/2 {
		max := 0
		for i := 1; i < length; i++ {
			if prices[i] > prices[i-1] {
				max += prices[i] - prices[i-1]
			}
		}
		return max
	}

	t := make([][]int, k)
	for i := 0; i < k; i++ {
		t[i] = make([]int, 2)
		t[i][0] = math.MinInt64
	}
	for _, p := range prices {
		if -p > t[0][0] {
			t[0][0] = -p
		}
		tp := t[0][0] + p
		if tp > t[0][1] {
			t[0][1] = tp
		}

		for i := 1; i < k; i++ {
			it0 := t[i-1][1] - p
			if it0 > t[i][0] {
				t[i][0] = it0
			}
			it1 := t[i][0] + p
			if it1 > t[i][1] {
				t[i][1] = it1
			}
		}
	}
	return t[k-1][1]
}
