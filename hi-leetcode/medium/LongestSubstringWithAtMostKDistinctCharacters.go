package medium

import "math"

/**
 * Given a string, find the length of the longest substring T that contains at most k distinct characters.
 *
 * For example, Given s = “eceba” and k = 2,
 * T is "ece" which its length is 3.
 */
/**
 * 题意：最多有K个不同字符的最长子串
 * 和 LengthOfLongestSubstringTwoDistinct1 类似
 */
func LengthOfLongestSubstringKDistinct2(s string, k int) int {
	var res, left, m = 0, 0, make(map[rune]int)
	for i, ci := range s {
		m[ci] = i
		for len(m) > k {
			cl := rune(s[left])
			if m[cl] == left {
				delete(m, cl)
			}
			left++
		}
		res = int(math.Max(float64(res), float64(i-left+1)))
	}
	return res
}
