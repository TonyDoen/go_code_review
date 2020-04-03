package algorithm

/**
 * 030-连续子数组的最大和
 *
 * 给一个数组,返回它的最大连续子序列的和？(子向量的长度至少是1)
 * 例如:{6,-3,-2,7,-15,1,2,2},连续子向量的最大和为8(从第0个开始,到第3个为止).
 *
 * 思路：
 * 令 函数 f(i)表示以第 i 个数字结尾的子数组的和,那我们这要求出 max(f(i)), i=1,2,...,n
 *        | arr[i]              i=0 or f(i-1)<=0
 * f(i) = {
 *        | arr[i] + f(i-1)    i!=0 or f(i-1)>0
 *
 * 从动态规划的状态转移方程来看,本来需要维护一张二维表记录每个阶段的最大和,再取最大值即为所求问题的解.
 * 但是 f(i) 只与 f(i−1) 状态有关,那么用一个临时变量记录 f(i−1), 另一个变量记录最大值,就可将空间复杂度降为常数级.
 *
 * 时间复杂度O(n),空间复杂度O(1)
 */
func FindGreatestSum(arr []int) int {
	if nil == arr {
		return -1
	}
	length, sum, max := len(arr), arr[0], arr[0]
	for i := 1; i < length; i++ {
		if sum <= 0 { // 如果sum <= 0,说明位置i之前的元素之和对后面的元素产生负影响或者没有影响,那么就需要抛弃之前的连续子序列,
			sum = arr[i]
		} else { // 否则,保留之前的连续子序列以及其和,接着向后遍历
			sum += arr[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}

/**
 * 052-正则表达式匹配
 *
 * 给你一个字符串 s 和一个字符规律 p,请你来实现一个支持 '.' 和'*' 的正则表达式匹配.
 * 1. '.' 匹配任意单个字符
 * 2. '*' 匹配零个或多个前面的那一个元素
 * 所谓匹配,是要涵盖 整个 字符串 s 的,而不是部分字符串.
 *
 * 说明:
 * s 可能为空,且只包含从a-z 的小写字母.
 * p 可能为空,且只包含从a-z 的小写字母,以及字符 .和 *
 *
 * 示例 1:
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串.
 *
 * 示例 2:
 * 输入:
 * s = "aa"
 * p = "a*"
 * 输出: true
 * 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'.因此,字符串 "aa" 可被视为 'a' 重复了一次.
 *
 * 示例 3:
 * 输入:
 * s = "aab"
 * p = "c*a*b"
 * 输出: true
 * 解释: 因为 '*' 表示零个或多个,这里 'c' 为 0 个, 'a' 被重复一次.因此可以匹配字符串 "aab".
 *
 *
 * 分析：
 * 题目的难点其实是在于 * 上面,如果没有这个 *,题目会变得非常简单,这里说一下题目的两个隐含条件：
 * 1. 一个就是 * 不会出现在字符串的开头
 * 2. 另外一个是 * 前面不能是 *,比如 “a * * b” 就不行
 */
/**
 * 递归方式的暴力深度优先搜索求解方法往往是搜索问题的万金油
 * 你只需要简单的考虑两件事情,
 * 1. 是这个问题是否可以划分为子问题,
 * 2. 是每个子问题有几种状态,就是在当前考虑的问题下,一共有多少种可能性.
 *
 * 字符串是可以划分成为一个个字符的,这样字符串比较的问题就会变成字符的比较问题
 * 决定 s[i,…n] 是否能够匹配 p[j,…m] 的条件
 * 是子问题 s[i+1,…n] 能不能够匹配 p[j+1,…m],另外还要看 s[i] 和 p[j] 是否匹配,
 *
 * 对于字符串 s 来说,没有特殊字符,当前问题中字符只会是字母,
 * 但是对于 p 来说,我们需要考虑两个特殊符号,还有字母,这里列举所有的可能,如果说当前的子问题是 s[i,…n] 和 p[j…m]:
 *
 * p[j] == s[i],               子问题成立与否取决于子问题 s[i+1,…n]  和 p[j+1,…m]
 * p[j] == ‘.’,                子问题成立与否取决于子问题 s[i+1,…n]  和 p[j+1,…m]
 * p[j+1] == ‘*’,s[i] != p[j],子问题成立与否取决于子问题 s[i,…n]    和 p[j+2,…m]
 * p[j+1] == ‘*’,s[i] == p[j],子问题成立与否取决于子问题 s[i+1,…n]  和 p[j,…m]
 */
func IsMatch(s, p string) bool {
	if s == p {
		return true
	}
	firstMatch := false
	if len(s) > 0 && len(p) > 0 && (s[0] == p[0] || '.' == p[0]) {
		firstMatch = true
	}

	if len(p) >= 2 && '*' == p[1] {
		return IsMatch(s, p[2:]) || (firstMatch && IsMatch(s[1:], p))
	}
	return firstMatch && IsMatch(s[1:], p[1:])
}

/**
 * url: https://www.cxyxiaowu.com/530.html
 *
 * 字符串匹配类动态规划的一些注意事项：
 * 1. 注意考虑是否需要考虑空串的情况,如果是的话,一般 DP 数组需要多开一格
 * 2. 在考虑递推方程前,确定子问题的区间和遍历方向
 * 3. 在思考递推方程的时候,重点思考当前子问题怎么变成之前求解过的子问题
 */
func IsMatch1(s, p string) bool {
	if s == p {
		return true
	}
	sLength, pLength := len(s), len(p)
	// dp[i][j] => is s[0, i - 1] match p[0, j - 1] ?
	dp := make([][]bool, sLength+1)
	for i := 0; i < sLength+1; i++ {
		dp[i] = make([]bool, pLength+1)
	}

	dp[0][0] = true // 两个空串是匹配
	for i := 1; i <= pLength; i++ {
		if '*' == p[i-1] {
			dp[0][i] = dp[0][i-2]
		} else {
			dp[0][i] = false
		}
	}

	for i := 1; i <= sLength; i++ {
		for j := 1; j <= pLength; j++ {
			if s[i-1] == p[j-1] || '.' == p[j-1] { // 看 s[0,...i-1] 和 p[0,...j-1]
				dp[i][j] = dp[i-1][j-1]
			}

			if '*' == p[j-1] {
				dp[i][j] = dp[i][j-2]                  // 看 s[0,...i] 和 p[0,...j-2]
				if p[j-2] == s[i-1] || '.' == p[j-2] { // 看 s[0,...i-1] 和 p[0,...j]
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[sLength][pLength]
}

/**
 * 011-二进制中1的个数
 *
 * 思路：
 * 如果一个整数不为0,那么这个整数至少有一位是1.
 * 如果我们把这个整数减1,那么原来处在整数最右边的1就会变为0,原来在1后面的所有的0都会变成1
 * (如果最右边的1后面还有0的话).其余所有位将不会受到影响.
 *
 * 例子：
 * 一个二进制数1100,从右边数起第三位是处于最右边的一个1.
 * 减去1后,第三位变成0,它后面的两位0变成了1,而前面的1保持不变,因此得到的结果是1011.
 *      1100 - 1
 *   => 1011
 * 我们发现减1的结果是把最右边的一个1开始的所有位都取反了.这个时候如果我们再把原来的整数和减去1之后的结果做与运算,从原来整数最右边一个1那一位开始所有位都会变成0.
 * 如1100&1011=1000.
 * 也就是说,把一个整数减去1,再和原整数做与运算,会把该整数最右边一个1变成0.
 * 那么一个整数的二进制有多少个1,就可以进行多少次这样的操作.
 *
 */
func NumberOf1InInt(n int) int {
	result := 0
	for 0 != n {
		result++
		n &= n - 1
	}
	return result
}

/**
 * 给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
 * 保证base和exponent不同时为0。
 *
 * 思路一：
 * 单层循环,时间复杂度O(abs(exponent))。
 *
 * 思路二：
 * 利用动态规划思想,递归实现。时间复杂度O(logN)。
 *       | a^(n/2) * a^(n/2)      n是偶数
 * a^n = {
 *       | a^(n/2) * a^(n/2) * a  n是奇数
 */
func Power(base float64, exponent int) float64 {
	if 0.0 == base && 0 == exponent {
		return 1
	}
	if 0.0 == base {
		return 0.0
	}
	reciprocal := false
	if exponent < 0 {
		reciprocal = true
		exponent = -exponent
	}

	result := 1.0
	for 0 != exponent {
		if 1 == exponent%2 {
			result *= base
		}
		base *= base
		exponent /= 2
	}

	if reciprocal {
		return 1 / result
	} else {
		return result
	}
}

/**
 * 040-数组中只出现一次的数字
 * 一个整型数组里除了两个数字之外,其他的数字都出现了两次。请写程序找出这两个只出现一次的数字
 *
 * 思路一：
 * HashMap。
 *
 */
func Find2SingleNumber1(arr []int) []int {
	if nil == arr || len(arr) < 1 {
		return nil
	}
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v] += 1
	}
	result := make([]int, 2)
	i := 0
	for k, v := range mp {
		if 1 == v {
			result[i] = k
			i++
		}
	}
	return result
}

/**
 * 思路二：
 * 异或法(相同的数字异或为0)
 * 1. 找到不同2个数异或差异位
 * 2. 根据异或差异位分别异或两组数, 找到2个只出现一次的数字
 */
func Find2SingleNumber2(arr []int) []int {
	if nil == arr || len(arr) < 1 {
		return nil
	}
	rex := 0
	// 从头到尾依次异或数组中的每一个数字，那么最终的结果刚好是那2个只出现一次的数字
	for _, v := range arr {
		rex ^= v
	}
	// find first bit
	idx := 0
	for (rex&1 == 0) && idx < 32 {
		rex >>= 1
		idx++
	}
	//
	result := make([]int, 2)
	for _, v := range arr {
		move := v >> uint(idx)
		if 1 == (move & 1) {
			result[0] ^= v
		} else {
			result[1] ^= v
		}
	}
	return result
}
