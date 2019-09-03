package easy

import (
	"container/list"
)

/**
 * Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.
 *
 * Examples:
 * Input: S = "a1b2"
 * Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * Input: S = "3z4"
 * Output: ["3z4", "3Z4"]
 *
 * Input: S = "12345"
 * Output: ["12345"]
 * Note:
 *
 * S will be a string with length at most 12.
 * S will consist only of letters or digits.
 */

/**
 * 题意： 字母大小写全排列
 * 这道题给了我们一个只包含字母和数字的字符串，让我们将字母以大小写进行全排列，给的例子很好的说明了题意。博主认为这道题给Easy有点不合适，至少应该是Medium的水准。这题主要参考了官方解答贴的解法，我们关心的是字母，数字的处理很简单，直接加上就可以了。比如说S = "abc"，那么先让 res = [""]，然后res中的每个字符串分别加上第一个字符a和A，得到 ["a", "A"]，然后res中的每个字符串分别加上第二个字符b和B，得到 ["ab", "Ab", "aB", "AB"]，然后res中的每个字符串分别加上第三个字符c和C，得到 ["abc", "Abc", "aBc", "ABc", "abC", "AbC", "aBC", "ABC"]，参见代码如下：
 */

func LetterCasePermutation4(s string) *list.List {
	res := list.New()
	res.PushBack(s)

	length := len(s)
	for i := 0; i < length; i++ {
		c := s[i]
		if '0' <= c && c <= '9' {
			continue
		}

		c ^= 32 // 因为我们知道 'A' = 65, 'B' = 66, 和 'a' = 97, 'b' = 98, 小写字母的ASCII码比大写字母多32，刚好是(1 << 5)
		// 顺序遍历
		for e := res.Front(); e != nil; e = e.Next() {
			tmp := []byte(e.Value.(string))
			tmp[i] = c
			res.PushFront(string(tmp))
		}
	}
	return res
}
