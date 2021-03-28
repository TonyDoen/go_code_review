package LCNOT

import (
	"bytes"
	"math"
)

/**
 * 编辑距离这个问题，
 * 它看起来十分困难，解法却出奇得简单漂亮，而且它是少有的比较实用的算法
 * （是的，我承认很多算法问题都不太实用）。
 *
 * 给定2个字符串 s1, s2, 计算将s1 转换成s2 所使用的最少操作数。
 * 可以对字符串进行如下3种操作
 * <1> 插入一个字符
 * <2> 删除一个字符
 * <3> 替换一个字符
 *
 *
 * 例子1：
 * input:  s1="horse", s2="ros"
 * output: 3
 * explain:
 *     horse -> rorse (将 'h' 替换为 'r')
 *     rorse -> rose  (将 'r' 删除)
 *     rose  -> ros   (将 'e' 删除)
 *
 * 例子2：
 * input:  s1="intention", s2="execution"
 * output: 5
 * explain:
 *     intention -> inention  (将 't' 删除)
 *     inention  -> enention  (将 'i' 替换为 'e')
 *     enention  -> exention  (将 'n' 替换为 'x')
 *     exention  -> exection  (将 'n' 替换为 'c')
 *     exection  -> execution (将 'u' 插入到 'c' 之后)
 *
 *
 * 思路：
 * 为什么说这个问题难呢，因为显而易见，它就是难，让人手足无措，望而生畏。
 *
 * 在日常生活中用到了这个算法。写错位了一段内容，我决定修改这部分内容让逻辑通顺。但是公众号文章最多只能修改 20 个字，且只支持增、删、替换操作（跟编辑距离问题一模一样），于是我就用算法求出了一个最优方案，只用了 16 步就完成了修改。
 * 高大上一点的应用，DNA 序列是由 A,G,C,T 组成的序列，可以类比成字符串。编辑距离可以衡量两个 DNA 序列的相似度，编辑距离越小，说明这两段 DNA 越相似。
 *
 *
 * 编辑距离问题就是给我们两个字符串s1和s2，只能用三种操作，让我们把s1变成s2，求最少的操作数。
 *
 * 前文[最长公共子序列]说过，解决两个字符串的动态规划问题，
 * 一般都是用两个指针i,j分别指向两个字符串的最后，然后一步步往前走，缩小问题的规模。
 *
 * 令 s1="rad"; s2="apple", 把 s1 变成 s2
 * init:
 *       s1 => r a d
 *                 i (指针)
 *       s2 => a p p l e
 *                     j (指针)
 *
 * step1: insert 'e'
 *       s1 => r a d e
 *                 i (指针)
 *       s2 => a p p l e
 *                   j (指针)
 *
 * step2: insert 'l'
 *       s1 => r a d l e
 *                 i (指针)
 *       s2 => a p p l e
 *                 j (指针)
 *
 * step3: insert 'p'
 *       s1 => r a d p l e
 *                 i (指针)
 *       s2 => a p p l e
 *                 j (指针)
 *
 * step4: replace 'd' -> 'p'
 *       s1 => r a p p l e
 *                 i (指针)
 *       s2 => a p p l e
 *               j (指针)
 *
 * step5: match and skip
 *       s1 => r a p p l e
 *               i (指针)
 *       s2 => a p p l e
 *             j (指针)
 *
 * step5: delete 'r'
 *       s1 => [r] a p p l e
 *             i (指针)
 *       s2 => a p p l e
 *             j (指针)
 *
 * 字符串操作需要5步
 *
 * 上面，就是j走完s2时，如果i还没走完s1，那么只能用删除操作把s1缩短为s2
 * 类似的，如果i走完s1时j还没走完了s2，那就只能用插入操作把s2剩下的字符全部插入s1。
 * 等会会看到，这两种情况就是算法的 base case。
 *
 * 伪代码：
 *   if s1[i] == s2[j]:
 *       skip
 *       i,j 同时前进
 *   else:
 *       三选一:
 *           insert (插入)
 *           delete (删除)
 *           replace(替换)
 *
 * 注意：
 *   这个「三选一」到底该怎么选择呢？很简单，全试一遍，哪个操作最后得到的编辑距离最小，就选谁。这里需要递归技巧，理解需要点技巧，先看下代码：
 *   def minDistance(s1, s2) -> int:
 *       def dp(i, j):
 *           # base case
 *           if i == -1: return j+1  # i 走完， 0~j+1 之间的字符删除
 *           if j == -1: return i+1  # j 走完， 0~i+1 之间的字符删除
 *
 *           if s1[i] == s2[j]:
 *               return dp(i-1, j-1) # skip
 *           else:
 *               return min(
 *                   dp(i, j-1) + 1,  # insert
 *                   dp(i-1, j) + 1,  # delete
 *                   dp(i-1, j-1)+1,  # replace
 *               )
 *
 *
 * 怎么能一眼看出存在重叠子问题呢？前文 动态规划之正则表达式 有提过，这里再简单提一下，需要抽象出本文算法的递归框架：
 *
 * def dp(i, j):
 *     dp(i - 1, j - 1) #1
 *     dp(i, j - 1)     #2
 *     dp(i - 1, j)     #3
 * 对于子问题dp(i-1,j-1)，如何通过原问题dp(i,j)得到呢？有不止一条路径，比如dp(i,j)->#1和dp(i,j)->#2->#3。一旦发现一条重复路径，就说明存在巨量重复路径，也就是重叠子问题。
 *
 *
 * 对于重叠子问题呢，前文 动态规划详解 介绍过，优化方法无非是备忘录或者 DP table。
 */

var (
	memo = make(map[string]*int)
	buff = new(bytes.Buffer)
)

func MinDistance(s1, s2 string) int {
	return helpMinDistance(s1, s2, len(s1)-1, len(s2)-1)
}

func helpMinDistance(s1, s2 string, i, j int) int {
	buff.Reset()
	_, _ = buff.WriteString(string(i))
	_, _ = buff.WriteString("->")
	_, _ = buff.WriteString(string(j))
	var key = buff.String()
	var mayVal = memo[key]
	if nil != mayVal {
		return *mayVal
	}

	// base case
	if -1 == i { // i 走完， 0~j+1 之间的字符删除
		return j + 1
	}
	if -1 == j { // j 走完， 0~i+1 之间的字符删除
		return i + 1
	}

	if s1[i] == s2[j] {
		var rs = helpMinDistance(s1, s2, i-1, j-1) // skip
		memo[key] = &rs
		return rs
	} else {
		var ins = helpMinDistance(s1, s2, i, j-1) + 1   // insert
		var del = helpMinDistance(s1, s2, i-1, j) + 1   // delete
		var rep = helpMinDistance(s1, s2, i-1, j-1) + 1 // replace
		var rs = min(ins, del, rep)
		memo[key] = &rs
		return rs
	}
}

func min(args ...int) int {
	var rs = math.MaxInt64
	for _, v := range args {
		if v < rs {
			rs = v
		}
	}
	return rs
}
