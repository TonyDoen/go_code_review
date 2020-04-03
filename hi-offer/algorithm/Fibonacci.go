package algorithm

/**
 * 007-斐波拉契数列
 *
 * 要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）。n<=39
 *
 * 0, 1, 1, 2, 3, 5, 8
 * 0  1  2  3  4  5  6
 *
 * 利用动态规划算法求解此问题：
 * 1. 阶段划分，状态表示，
 * 2. 状态转移方程：F(n) = F(n-1) + F(n-2)
 * 3. 边界条件：F(0) = 0, F(1) = 1.
 *
 * 解决此问题的三种方式：
 * 1. 递归-时间复杂度 O(2^n)
 * 2. 备忘录递归-时间复杂度 O(n)
 * 3. 自底向上迭代-时间复杂度 O(n)
 */
func Fibonacci1(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci1(n-1) + Fibonacci1(n-2)
}
func Fibonacci2(n int) int {
	if n < 0 {
		return n
	}
	memo := make([]int, n+1)
	for i := 1; i < n; i++ {
		memo[i] = -1
	}

	return helpFibonacci2(n, memo)
}
func helpFibonacci2(n int, memo []int) int {
	if -1 != memo[n] {
		return memo[n]
	}
	if n < 3 {
		memo[n] = 1
	} else {
		memo[n] = helpFibonacci2(n-1, memo) + helpFibonacci2(n-2, memo)
	}
	return memo[n]
}

func Fibonacci3(n int) int {
	if n < 2 {
		return n
	}

	result, _1, _2 := 0, 0, 1
	for i := 1; i < n; i++ {
		result = _1 + _2
		_1 = _2
		_2 = result
	}
	return result
}

/**
 * 008-跳台阶
 * 一只青蛙一次可以跳上1级台阶，也可以跳上3级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
 */
func Jump(n int) int {
	if n < 1 {
		return -1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-3]
	}
	return dp[n]
}

/**
 * 009-变态跳台阶
 * 一只青蛙一次可以跳上1级台阶，也可以跳上3级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
 *
 * 这里主要有两种思路，我感觉第二种更好理解一点。
 * 1. 假设：f(n)表示：n个台阶第一次1,2,...n阶的跳法数; 若第一次跳了1阶，则还剩n-1阶，
 * 　　假设：f(n-1)表示：n-1个台阶第一次1,2,...n-1阶的跳法数; 若第一次跳了2阶，则还剩n-2阶，
 * 　　假设：f(n-2)表示：n-1个台阶第一次1,2,...n-2阶的跳法数;
 *     ...
 * 　　把所以可能的情况（第一次可能跳1,2,...,n阶）加起来：
 * 　　可以求出：f(n) = f(n-1) + f(n-2) + ... + f(1)
 * 　　递归：f(n-1) = f(n-2) + ... + f(1)
 * 　　可以求出：f(n) = 2*f(n-1)
 *
 * 2. 每个台阶可以跳或者不跳
 *   Sum(C(x, n-1)) = 2^(n-1)
 */
func FreakJump1(n int) int {
	if n < 3 {
		return n
	}
	return 2 * FreakJump1(n-1)
}
func FreakJump2(n int) int {
	if n < 3 {
		return n
	}
	result := 1
	for i := 1; i < n; i++ {
		result *= 2
	}
	return result
}

/**
 * 我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。
 * 请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
 *  _               _____________________
 * | |             |                     |
 * |_|             |_____________________|
 * 2*1的小矩形       2*n的大矩形
 *
 * 动态规划算法求解此问题：
 * 1. 阶段划分，状态表示，
 * 2. 状态转移方程：F(n) = F(n-1) + F(n-2)
 * 3. 边界条件：F(0) = 0, F(1) = 1, F(2) = 2
 *
 * 状态转移方程的含义为 ：
 * 要覆盖n个2*1  有f(n)种，
 * 1. 第一种是从n-1个2*1的f(n-1)最后一格竖着覆盖一种方式，
 * 2. 第二种是从n-2个2*1的f(n-2)最后两格横着覆盖一种方式得到的。
 *
 * 思路同Fibonacci数列
 */
func RectCover(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
