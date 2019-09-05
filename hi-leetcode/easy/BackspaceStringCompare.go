package easy

/**
 * Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
 *
 * Example 1:
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 *
 * Explanation: Both S and T become "ac".
 *
 * Example 2:
 * Input: S = "ab##", T = "c#d#"
 * Output: true
 *
 * Explanation: Both S and T become "".
 *
 * Example 3:
 * Input: S = "a##c", T = "#a#c"
 * Output: true
 *
 * Explanation: Both S and T become "c".
 *
 * Example 4:
 * Input: S = "a#c", T = "b"
 * Output: false
 *
 * Explanation: S becomes "c" while T becomes "b".
 *
 * Note:
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S and T only contain lowercase letters and '#' characters.
 * Follow up:
 *
 * Can you solve it in O(N) time and O(1)space?
 */

/**
 * 退格字符串比较
 * 这道题给了我们两个字符串，里面可能会有井号符#，这个表示退格符，键盘上的退格键我们应该都很熟悉吧，当字打错了的时候，肯定要点退格键来删除的。当然也可以连续点好几下退格键，这样就可以连续删除了，在例子2和3中，也确实能看到连续的井号符。
 */

func BackspaceCompare1(s, t string) bool {
	sBuf, sLen := helper(s)
	tBuf, tLen := helper(t)
	if sLen != tLen {
		return false
	}
	for i := 0; i < sLen; i++ {
		if sBuf[i] != tBuf[i] {
			return false
		}
	}
	return true
}

func helper(str string) ([]byte, int) {
	length := len(str)
	tmp := 0
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		c := str[i]
		if '#' == c {
			if tmp > 0 {
				tmp--
			}
		} else {
			buf[tmp] = c
			tmp++
		}
	}
	return buf, tmp
}

func BackspaceCompare2(s, t string) bool {
	sLen := len(s) - 1
	tLen := len(t) - 1
	sCnt := 0
	tCnt := 0
	for sLen >= 0 && tLen >= 0 {
		for sLen >= 0 && ('#' == s[sLen] || sCnt > 0) {
			if '#' == s[sLen] {
				sCnt++
			} else {
				sCnt--
			}
			sLen--
		}
		for tLen >= 0 && ('#' == t[tLen] || tCnt > 0) {
			if '#' == t[tLen] {
				tCnt++
			} else {
				tCnt--
			}
			tLen--
		}
		if sLen < 0 || tLen < 0 {
			return sLen == tLen
		}
		if s[sLen] != t[tLen] {
			return false
		} else {
			sLen--
			tLen--
		}
	}
	return sLen == tLen
}
