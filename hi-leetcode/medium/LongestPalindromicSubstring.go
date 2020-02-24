package medium

/**
 * Longest Palindromic Substring (最长回文子串)
 *
 * 示例 1：
 *   输入: "babad"
 *   输出: "bab"
 *   注意: "aba" 也是一个有效答案。
 *
 * 示例 2：
 *   输入: "cbbd"
 *   输出: "bb"
 */

/**
     * 暴力解法为遍历所有子串，逐个判断是否是回文字串。
     * <p>
     * 解释1:
     * 在动态规划的思想中，总是希望把问题划分成相关联的子问题；然后从最基本的子问题出发来推导较大的子问题，直到所有的子问题都解决。
     * 假设字符串s的长度为length，建立一个 length*length 的矩阵dp。
     * 令 dp[i][j] 表示 S[i] 至 S[j] 所表示的子串是否是回文子串。
     * <p>
     * 1. 当 i == j，dp[i][j] 是回文子串（单字符都是回文子串）；
     * 2. 当 j - i < 3，只要 S[i] == S[j]，则 dp[i][j] 是回文子串（如"aa"，"aba"），否则不是；
     * 3. 当 j - i >= 3，如果 S[i] == S[j] && dp[i+1][j-1] ，则 dp[i][j] 是回文子串，否则不是 。
     * <p>
     * 由此可以写出状态转移方程：
     *           ⎪ true,                            i == j
     * dp[i][j]= ⎨ S[i] == S[j],                    j-i < 3
     *           ⎪ S[i] == S[j] && dp[i+1][j-1],    j-i >= 3
     * <p>
     * 时间复杂度：O(n²)。
     * 空间复杂度：O(n²)。
**/
func LongestPalindromicSubstring1(src string) string {
	length := len(src)
	result := string(src[0])
	// 2-D array initialization
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}

	for gap := 0; gap < length; gap++ {
		for i := 0; i < length-gap; i++ {
			j := i + gap
			if src[i] == src[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j+1-i > len(result) {
					result = src[i : j+1]
				}
			}
		}
	}
	return result
}

/**
 * 中心扩展算法
 * 中心扩展就是把给定的字符串的每一个字母或两个字母之间空隙当做中心，向两边扩展，这样来找
 * 1. 长度为奇数的回文串，比如a, aba, abcba，以字母为中心
 * 2. 长度为偶数的回文串，比如aa, abba，以两个字母之间空隙为中心
 * <p>
 * 时间复杂度：O(n²)。
 * 空间复杂度：O(1)。
 */
func LongestPalindromicSubstring2(src string) string {
	start, end, length := 0, 0, len(src) // 记录回文子串的开始位置
	for i := 0; i < length; i++ {
		// 以每个字符为中心去扩展，例如"aba"就是以'b'为中心
		len1 := expandAroundCenter(src, i, i)
		// 以两字母之间为中心去扩展，例如 "abba" 的中心在两个 'b'之间
		len2 := expandAroundCenter(src, i, i+1)
		len3 := max(len1, len2)
		if len3 > end-start {
			start = i - (len3-1)/2
			end = i + len3/2
		}
	}
	return src[start : end+1]
}
func expandAroundCenter(src string, left, right int) int {
	length := len(src)
	for ; left >= 0 && right < length && src[left] == src[right]; {
		left--
		right++
	}
	return right - left - 1
}

/**
 * Manacher（马拉车）算法
 * Manacher算法，又叫"马拉车"算法，可以在时间复杂度为O(n)的情况下求解一个字符串的最长回文子串长度的问题。
 * 1. 将初始字符串每个字符左右两边填充’#’(也可以是其它字符)，巧妙地解决对称数量奇偶的问题（如"aba"变成"#a#b#a#","bb"变成"#b#b#",处理后的回文子串都是奇数）；
 * 2. 遍历整个字符串，用一个数组来记录以该字符为中心的回文子串半径，并记录已经扩展到的右边界；
 * 3. 每一次遍历的时候，如果该字符在已知回文串最右边界的覆盖下，那么就计算其相对最右边界回文串中心对称的位置，得出已知回文串的长度；
 * 4. 判断该长度和右边界，如果达到了右边界，那么需要继续进行中心扩展探索。当然，如果第3步该字符没有在最右边界的"羽翼"下，则直接进行中心扩展探索。进行中心扩展探索的时候，同时又更新右边界；
 * 5. 最后得到最长回文子串之后，去掉其中的特殊符号即可。
 * <p>
 * 时间复杂度：O(n)，这个算法在循环的时候，要么扩展右边界，要么直接得出结论，时间复杂度可以到O(n)。
 * 空间复杂度：O(n)。
 */
func LongestPalindromicSubstring3(src string) string {

	// 1. 将初始字符串每个字符左右两边填充’#’(也可以是其它字符)
	tmp := "#"
	for _, c := range src {
		tmp = tmp + string(c) + "#"
	}

	// 2. 遍历整个字符串，用一个数组来记录以该字符为中心的回文子串半径，并记录已经扩展到的右边界；
	length := len(tmp)
	maxRight, maxLength, pos, halfLenArr := 0, 0, 0, make([]int, length) // 记录 最长半径范围, 最长回文字符串长度, 当前最长半径的对称轴, 每个位置最长回文长度
	for i := 0; i < length; i++ {
		if i < maxRight { // 当前 字符 在最大半径范围 左边
			if halfLenArr[2*pos-i] > maxRight-i {
				halfLenArr[i] = maxRight - i
			} else {
				halfLenArr[i] = halfLenArr[2*pos-i]
			}
		} else { // 当前 字符 在 最大半径 范围 右边(没有被遍历过)
			halfLenArr[i] = 1
		}

		// 在先前 计算的 回文长度 基础 上 扩展遍历
		for ; i-halfLenArr[i] >= 0 && i+halfLenArr[i] < length && tmp[i-halfLenArr[i]] == tmp[i+halfLenArr[i]]; {
			halfLenArr[i] += 1
		}
		if halfLenArr[i]+i-1 > maxRight {
			maxRight = halfLenArr[i] + i - 1
			pos = i
		}
		if maxLength < halfLenArr[i] {
			maxLength = halfLenArr[i]
		}
	}

	// 5. 最后得到最长回文子串之后，去掉其中的特殊符号即可。去掉之前添加的#
	start := (pos+1)/2 - maxLength/2
	end := start + maxLength - 1

	return src[start:end]
}
