package medium

import "math"

/**
 * Given a string s , find the length of the longest substring t  that contains at most 2 distinct characters.
 *
 * Example 1:
 * Input: "eceba"
 * Output: 3
 * Explanation: tis "ece" which its length is 3.
 *
 * Example 2:
 * Input: "ccaabbb"
 * Output: 5
 * Explanation: tis "aabbb" which its length is 5.
 */

/**
 * 题意：最多有两个不同字符的最长子串
 * 这道题给我们一个字符串，让我们求最多有两个不同字符的最长子串。
 * 用 HashMap 来映射每个字符最新的坐标，比如题目中的例子 "eceba"，遇到第一个e，映射其坐标0，遇到c，映射其坐标1，遇到第二个e时，映射其坐标2，当遇到b时，映射其坐标3，每次我们都判断当前 HashMap 中的映射数，如果大于2的时候，那么需要删掉一个映射，我们还是从 left=0 时开始向右找，看每个字符在 HashMap 中的映射值是否等于当前坐标 left，比如第一个e，HashMap 此时映射值为2，不等于 left 的0，那么 left 自增1，遇到c的时候，HashMap 中c的映射值是1，和此时的 left 相同，那么我们把c删掉，left 自增1，再更新结果，以此类推直至遍历完整个字符串
 */

func LengthOfLongestSubstringTwoDistinct1(s string) int {
	var res, left int
	m := make(map[byte]int)
	length := len(s)
	for i := 0; i < length; i++ {
		m[s[i]] = i
		for len(m) > 2 {
			if m[s[left]] == left {
				delete(m, s[left])
			}
			left++
		}
		res = int(math.Max(float64(res), float64(i-left+1)))
	}
	return res
}

/**
 * 这里我们使用若干的变量，其中 cur 为当前最长子串的长度，first 和 second 为当前候选子串中的两个不同的字符，cntLast 为 second 字符的连续长度。我们遍历所有字符，假如遇到的字符是 first 和 second 中的任意一个，那么 cur 可以自增1，否则 cntLast 自增1，因为若是新字符的话，默认已经将 first 字符淘汰了，此时候选字符串由 second 字符和这个新字符构成，所以当前长度是 cntLast+1。然后再来更新 cntLast，假如当前字符等于 second 的话，cntLast 自增1，否则均重置为1，因为 cntLast 统计的就是 second 字符的连续长度。然后再来判断若当前字符不等于 second，则此时 first 赋值为 second， second 赋值为新字符。最后不要忘了用 cur 来更新结果 res
 */
func LengthOfLongestSubstringTwoDistinct2(s string) int {
	var res, cur, cntLast int
	var first, second byte
	length := len(s)
	for i := 0; i < length; i++ {
		c := s[i]
		if c == first || c == second {
			cur = cur + 1
		} else {
			cur = cntLast + 1
		}
		if c == second {
			cntLast = cntLast + 1
		} else {
			cntLast = 1
		}
		if c != second {
			first = second
			second = c
		}
		res = int(math.Max(float64(res), float64(cur)))
	}
	return res
}
