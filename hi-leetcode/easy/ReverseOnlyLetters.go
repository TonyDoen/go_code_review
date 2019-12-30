package easy

/**
 * Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.
 *
 * Example 1:
 * Input: "ab-cd"
 * Output: "dc-ba"
 *
 * Example 2:
 * Input: "a-bC-dEf-ghIj"
 * Output: "j-Ih-gfE-dCba"
 *
 * Example 3:
 * Input: "Test1ng-Leet=code-Q!"
 * Output: "Qedo1ct-eeLg=ntse-T!"
 * Note:
 *
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122
 * S doesn't contain \ or "
 */

/**
 * 题意：只翻转字母
 * 思路：这道题给了一个由字母和其他字符组成的字符串，让我们只翻转其中的字母，并不是一道难题，解题思路也比较直接。可以先反向遍历一遍字符串，只要遇到字母就直接存入到一个新的字符串 res，这样就实现了对所有字母的翻转。但目前的 res 中就只有字母，还需要将原字符串S中的所有的非字母字符加入到正确的位置，可以再正向遍历一遍原字符串S，遇到字母就跳过，否则就把非字母字符加入到 res 中对应的位置，
 */
func ReverseOnlyLetters1(src string) string {
	length := len(src)
	res := make([]byte, length)
	for i, j := length - 1, 0; i >= 0; i-- {
		c := src[i]
		if isLetter(c) {
			res[j] = c
			j++
		}
	}
	for i := 0; i < length; i++ {
		c := src[i]
		if !isLetter(c) {
			for j := length - 2; j >= i; j-- {
				res[j+1] = res[j]
			}
			res[i] = c
		}
	}
	return string(res)
}

func ReverseOnlyLetters2(src string) string {
	length := len(src)
	i, j, res := 0, length-1, make([]byte, length)
	for ; i < j; {
		ic, jc := src[i], src[j]
		if !isLetter(ic) {
			res[i] = ic
			i++
		} else if !isLetter(jc) {
			res[j] = jc
		    j--
		} else {
		    res[i] = jc
		    res[j] = ic
		    i++
		    j--
		}
	}
	return string(res)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}