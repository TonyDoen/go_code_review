package medium

import "bytes"

/**
 * We have a string S of lowercase letters, and an integer array shifts.
 * Call the shift of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a').
 * For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.
 * Now for each shifts[i] = x, we want to shift the first i+1 letters of S, x times.
 * Return the final string after all such shifts to S are applied.
 *
 * Example 1:
 * Input: S = "abc", shifts = [3,5,9]
 * Output: "rpl"
 *
 * Explanation:
 * We start with "abc".
 * After shifting the first 1 letters of S by 3, we have "dbc".
 * After shifting the first 2 letters of S by 5, we have "igc".
 * After shifting the first 3 letters of S by 9, we have "rpl", the answer.
 *
 * Note:
 * 1 <= S.length = shifts.length <= 20000
 * 0 <= shifts[i] <= 10 ^ 9
 */

/**
 * 题意
 * 漂移字母
 *
 * 这道题让我们对字母进行漂移，给了一个shifts数组，里面是对应对需要漂移值，但是需要注意的是，shifts[i] 表示对于原字符串 [0, i] 范围内的所有的字符都进行 shifts[i] 的漂移，那么实际上第一个字母其实把shifts数组所有的数字都漂移了一遍，而第二个字母则是把shifts数组从第二个数字开始到最后的所有数字都漂移了，而最后一个字母就只漂移了最后一个数字。这不就是一个反向累加和数组么，只要建立了反向累加和数组，那么每个位子上的数字就是对应的字母的漂移值了。为了节省空间，我们就不另建数组了，直接在shifts数组上累加就行了，注意我们的累加值要对26取余，因为累加和数组可能会整型溢出，取余后就不会有这个问题，而且由于字母漂移26次后，都会回到原来的位置，所以对26取余并不会影响到最后的结果。
 * 反向累加和数组建立好了之后，就要开始对字母进行漂移了，这里还有个需要注意的地方，我们不能直接用原字母加上漂移值，因为一旦超过了 'z' 的时候，是需要从 'a' 重新的开始的，为了处理所有的情况，我们可以使用一个很常用的trick，就是先算出字母到 'a' 之间的距离，然后加上漂移值，再对26取余，这就是漂移后与 'a' 的距离了，再加上 'a' 变成字母即可
 */

func ShiftingLetters1(src string, shifts []int) string {
	length := len(shifts)
	for i := length - 2; i > -1; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}
	buff := new(bytes.Buffer)
	for i := 0; i < length; i++ {
		c := (src[i] - 'a' + uint8(shifts[i])) % 26 + 'a'
		buff.WriteByte(c)
	}
	return buff.String()
}
