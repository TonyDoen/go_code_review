package medium

/**
 * Alice has a hand of cards, given as an array of integers.
 * Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.
 * Return true if and only if she can.
 *
 * Example 1:
 * Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
 * Output: true
 *
 * Explanation:
 * Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
 *
 * Example 2:
 * Input: hand = [1,2,3,4,5], W = 4
 * Output: false
 *
 * Explanation:
 * Alice's hand can't be rearranged into groups of 4
 *
 * Note:
 * 1 <= hand.length <= 10000
 * 0 <= hand[i] <= 10^9
 * 1 <= W <= hand.length
 */

/**
 * 一手顺子牌
 * 这里给了我们一个W，规定了顺子的最小长度，那么我们就拿例子1来模拟下打牌吧，首先摸到了牌之后，肯定要先整牌，按从小到大的顺序排列，这里我们就不考虑啥3最大，4最小啥的，就统一按原始数字排列吧：1 2 2 3 3 4 6 7 8
 * 下面要来组顺子，既然这里是3张可连，那么我们从最小的开始连呗。其实这道题还是简化了许多，真正打牌的时候，即便是3张起连，那么连4张5张都是可以的，可以这里限定了只能连W张，就使得题目变简单了。我们用贪婪算法就可以了，首先从1开始，那么一定得有2和3，才能起连，若少了任何一个，都可以直接返回false，好那么取出这三张后
 */

func IsStraightHand1(hand []int, w int) bool {
	mp := make(map[int]int)
	start := hand[0]
	for _, i := range hand {
		cnt := mp[i]
		if 0 == cnt {
			cnt = 1
		} else {
			cnt += 1
		}
		mp[i] = cnt

		if start > i {
			start = i
		}
	}

	for len(mp) > 0 {
		for i := 0; i < w; i++ {
			want := start + i
			v := mp[want]
			if v <= 0 {
				return false
			}
			v--
			if 0 == v {
				delete(mp, want)
			} else {
				mp[want] = v
			}
		}
		for k := range mp {
			start = k
			break
		}
	}
	return true
}
