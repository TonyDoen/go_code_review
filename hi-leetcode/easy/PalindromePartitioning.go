package easy

import "container/list"

/**
 * url: https://blog.csdn.net/zangdaiyang1991/article/details/89338652
 * url: https://blog.csdn.net/sinat_27908213/article/details/80599460
 *
 * palindrome-partitioning
 * Given a string s, partition s such that every substring of the partition is a palindrome.
 * Return all possible palindrome partitioning of s.
 *
 * For example, given s ="aab",
 * Return
 *   [
 *     ["aa","b"],
 *     ["a","a","b"]
 *   ]
 *
 */
/**
 * 题意：给出所有可能的回文串分词
 * 思路：回溯算法实现（BackTracking）
 *
 * 回溯法有通用解法的美称，对于很多问题，如迷宫等都有很好的效果。
 * 回溯算法实际上一个类似枚举的深度优先搜索尝试过程，主要是在搜索尝试过程中寻找问题的解，当发现已不满足求解条件时，就“回溯”返回（也就是递归返回），尝试别的路径。
 * 回溯法说白了就是穷举法。回溯法一般用递归来实现。
 *
 * 回溯法一般都用在要给出多个可以实现最终条件的解的最终形式。
 * 回溯法要求对解要添加一些约束条件。
 * 总的来说，如果要解决一个回溯法的问题，通常要确定三个元素：
 * 1、选择。对于每个特定的解，肯定是由一步步构建而来的，而每一步怎么构建，肯定都是有限个选择，要怎么选择，这个要知道；同时，在编程时候要定下，优先或合法的每一步选择的顺序，一般是通过多个if或者for循环来排列。
 * 2、条件。对于每个特定的解的某一步，他必然要符合某个解要求符合的条件，如果不符合条件，就要回溯，其实回溯也就是递归调用的返回。
 * 3、结束。当到达一个特定结束条件时候，就认为这个一步步构建的解是符合要求的解了。把解存下来或者打印出来。对于这一步来说，有时候也可以另外写一个issolution函数来进行判断。注意，当到达第三步后，有时候还需要构建一个数据结构，把符合要求的解存起来，便于当得到所有解后，把解空间输出来。这个数据结构必须是全局的，作为参数之一传递给递归函数。
 *
 *
 */
func PalindromePartitioning(src string) *list.List {
	result := list.New()
	helpPalindromePartitioning(result, list.New(), src)
	return result
}
func helpPalindromePartitioning(result, one *list.List, src string) {
	length := len(src)
	if length < 1 {
		// copy one
		tmp, tmpOne := one.Front(), list.New()
		for nil != tmp {
			tmpOne.PushBack(tmp.Value)
			tmp = tmp.Next()
		}
		result.PushBack(tmpOne)
		return
	}
	for i := 1; i <= length; i++ {
		sub := src[0:i]
		if isPalindrome(sub) {
			one.PushBack(sub)
			helpPalindromePartitioning(result, one, src[i:length])
			one.Remove(one.Back())
		}
	}
}
func isPalindrome(s string) bool {
	left, right := 0, len(s)
	for left < right {
		if s[left:left+1] != s[right-1:right] {
			return false
		}
		left++
		right--
	}
	return true
}

/**
 * Given a string s, partition s such that
 * every substring of the partition is a palindrome.
 *
 * Return the minimum cuts needed for a palindrome partitioning of s.
 *
 * For example, given s ="aab",
 * Return 1 since the palindrome partitioning ["aa","b"]
 * could be produced using 1 cut.
 */
/**
 * 思路1，动态规划(使用一维数组，复杂度高)
 * 动态规划问题(一维数组表示，复杂度大)
 * dp[i] - 表示子串（0，i）的最小回文切割，则最优解在dp[s.length-1]中。
 *
 * 分几种情况：
 * 1.初始化：当字串s.substring(0,i+1)(包括i位置的字符)是回文时，dp[i] = 0(表示不需要分割)；否则，dp[i] = i（表示至多分割i次）;
 * 2.对于任意大于1的i，如果s.substring(j,i+1)(j<=i,即遍历i之前的每个子串)是回文时，dp[i] = min(dp[i], dp[j-1]+1);
 * 3.如果s.substring(j,i+1)(j<=i)不是回文时，dp[i] = min(dp[i],dp[j-1]+i+1-j);
 *
 */
func MinCut(src string) int {
	length := len(src)
	if length < 1 {
		return 0
	}
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		str := src[0 : i+1]
		if isPalindrome(str) {
			dp[i] = 0
		} else {
			dp[i] = i
		}
		if 0 == dp[i] {
			continue
		}
		for j := 1; j <= i; j++ {
			subStr := src[j : i+1]
			if isPalindrome(subStr) {
				dp[i] = min(dp[i], dp[j-1]+1)
			} else {
				dp[i] = min(dp[i], dp[j-1]+i-j+1)
			}
		}
	}
	return dp[length-1]
}

/**
 * 动态规划的题，最主要就是写出状态转移方程
 * 状态转移，其实就是怎么把一个大的状态表示为两个或者多个已知的状态
 * 以此题为例，设f[i][j]为最小的切点数，那么有：
 * 1、s[i][j]为回文字符串，则f[i][j] = 0;
 * 2、s[i][j]不是回文字符串，增加一个切点p，将s[i][j]切割为两端s[i][p]、s[p+1][j],则f[i][j] = f[i][p]+f[p+1][j]+1
 *
 * 所谓的状态转移方程就是上面的式子
 * 接着来看看怎么组织程序，先看看状态转移的思路：
 * 以"aab"为例，"aab"明显不是回文串
 * 所以 f("aab") = min( (f("a")+f("ab")) , (f("aa")+f("b")) ) + 1;
 * f("a") = 0;
 * f("ab") = f("a")+f("b") +1  = 0+0+1 = 1;
 * f("aa") = 0;
 * f("b") = 0;
 * 即f("aab") = 1;
 */
func MinCut0(src string) int {
	length := len(src)
	if length < 1 {
		return 0
	}
	// init dp
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	for gap := 0; gap <= length; gap++ {
		for i, j := 0, gap; j < length; {
			if isPalindrome(src[i : j+1]) {
				dp[i][j] = 0
			} else {
				min := length - 1
				for p := i; p < j; p++ {
					a := dp[i][p]
					b := dp[p+1][j]
					a = a + b + 1
					if min > a {
						min = a
					}
				}
				dp[i][j] = min
			}

			i++
			j++
		}
	}
	return dp[0][length-1]
}