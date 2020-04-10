package medium

import "strings"

/**
 * url: https://leetcode.com/problems/push-dominoes/submissions/
 * url: https://leetcode.com/problems/push-dominoes/discuss/132332/C%2B%2BJavaPython-Two-Pointers
 *
 * There are N dominoes in a line, and we place each domino vertically upright.
 * In the beginning, we simultaneously push some of the dominoes either to the left or to the right.
 *
 * After each second, each domino that is falling to the left pushes the adjacent domino on the left.
 * Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.
 * When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.
 * For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.
 * Given a String "S" representing the initial state. S[i] = 'L', if the i-th domino has been pushed to the left; S[i] = 'R', if the i-th domino has been pushed to the right; S[i] = '.', if the i-th domino has not been pushed.
 * Return a String representing the final state.
 *
 * Example 1:
 * Input: ".L.R...LR..L.."
 * Output: "LL.RR.LLRRLL.."
 *
 * Example 2:
 * Input: "RR.L"
 * Output: "RR.L"
 *
 * Explanation: The first domino expends no additional force on the second domino.
 *
 * Note:
 * 0 <= N <= 10^5
 * String dominoes contains only 'L', 'R' and '.'
 *
 *
 * 题意：
 * 推多米诺骨牌
 *
 * 思路：
 * 这道题给我们摆好了一个多米诺骨牌阵列，但是与一般的玩法不同的是，这里没有从一头开始推，而是在很多不同的位置分别往两个方向推，结果是骨牌各自向不同的方向倒下了，而且有的骨牌由于左右两边受力均等，依然屹立不倒，这样的话骨牌就很难受了，能不能让哥安心的倒下去？！生而为骨牌，总是要倒下去啊，就像漫天飞舞的樱花，秒速五厘米的落下，回到最终归宿泥土里。喂，不要给骨牌强行加戏好么！～ 某个位置的骨牌会不会倒，并且朝哪个方向倒，是由左右两边受到的力的大小决定的
 *
 *  1）R....R  ->  RRRRRR
 * 这是当两个向右推的操作连在一起时，那么中间的骨牌毫无悬念的都要向右边倒去。
 *  2）L....L  ->  LLLLLL
 * 同理，当两个向左推的操作连在一起时，那么中间的骨牌毫无悬念的都要向左边倒去。
 *  3）L....R  ->  L....R
 * 当左边界的骨牌向左推，右边界的骨牌向右推，那么中间的骨牌不会收到力，所以依然保持坚挺。
 *  4）R....L  -> RRRLLL   or   R.....L  ->  RRR.LLL
 * 当左边界的骨牌向右推，右边界的骨牌向左推时，就要看中间的骨牌个数了，若是偶数，那么对半分，若是奇数，那么最中间的骨牌保持站立，其余的对半分。
 *
 * 由于上述四种情况包含了所有的情况，所以我们的目标就是在字符串中找出中间是‘点’的小区间，为了便于我们一次遍历就处理完，我们在dominoes字符串左边加个L，右边加个R，这并不会影响骨牌倒下的情况。我们使用双指针来遍历，其中i初始化为0，j初始化为1，当j指向‘点’时，我们就跳过，目标是i指向小区间的左边界，j指向右边界，然后用 j-i-1 算出中间‘点’的个数，为0表示中间没有点。若此时 i>0，则将左边界加入结果res中。若左右边界相同，那么中间的点都填成左边界，这是上述的情况一和二；若左边界是L，右边界是R，则是上述的情况三，中间还是保持点不变；若左边界是R，右边界是L，则是情况四，那么先加 mid/2 个R，再加 mid%2 个点，最后加 mid/2 个L即可。然后i更新为j，继续循环即可
 */
func PushDominoes(d string) string {
	d = "L" + d + "R"
	var result strings.Builder
	for i, j, length := 0, 1, len(d); j < length; j++ {
		ci, cj := d[i], d[j]
		if '.' == cj {
			continue
		}
		middle := j - i - 1
		if i > 0 {
			result.WriteByte(ci)
		}
		if ci == cj {
			for k := 0; k < middle; k++ {
				result.WriteByte(ci)
			}
		} else if 'L' == ci && 'R' == cj {
			for k := 0; k < middle; k++ {
				result.WriteByte('.')
			}
		} else {
			for k := 0; k < middle/2; k++ {
				result.WriteByte('R')
			}
			if 1 == middle%2 {
				result.WriteByte('.')
			}
			for k := 0; k < middle/2; k++ {
				result.WriteByte('L')
			}
		}
		i = j
	}
	return result.String()
}
