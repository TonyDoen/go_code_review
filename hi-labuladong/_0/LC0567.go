package _0

/**
  滑动窗口算法框架

  void slidingWindow (String s, String t) {
      Map<Character, Integer> need = new HashMap<>(), window = new HashMap<>();
      for (char c : t.toCharArray()) {
          need.put(c, need.getOrDefault(c, 0)+1);
      }

      int left = 0, right = 0, valid = 0;
      while (right < s.length()) {
          // c 是将要移入窗口的字符
          char c = s.charAt(right);
          // 右移窗口
          right++;
          // 进行窗口内数据的一系列更新
          // ...

          // debug
          System.out.println("window: {}...");

          // 判断左侧窗口是否收缩
          while (window needs shrink) {
              // d 是将要移出窗口的字符
              char d = s.charAt(left);
              // 左移窗口
              left++;
              // 进行窗口内数据的一系列更新
              // ...
          }
      }
  }
  时间复杂度是 O(N)
*/

/*
 * LeetCode 567 题， Permutation in String
 *
 * 给定2个字符串 s1, s2; 判断 s2 是否包含 s1 的排列
 * 即, 第一个字符串的排列之一是第二个字符串的子串
 *
 * 例子1：
 * input: s1 = "ab"; s2 = "eidbaooo"
 * output: true
 * explain: s2 包含 s1 的排列之一 ["ba"]
 *
 * 例子2：
 * input: s1 = "ab"; s2 = "eidboaooo"
 * output: false
 * explain:
 *
 * 注意： s1 是可以包含重复字符的
 *
 * 这种题目，是明显的滑动窗口
 * 给你一个S和T，请问S中是否存在一个子串，包含T中所有字符且不包含其他字符
 *
 */
func CheckInclusion(t, s string) bool {
	var need, window, tLength = make(map[uint8]int), make(map[uint8]int), len(t)
	for i := 0; i < tLength; i++ {
		need[t[i]]++
	}

	var left, right, valid = 0, 0, 0
	for sLength := len(s); right < sLength; {
		var c = s[right]
		right++

		// 进行窗口内数据的一系列更新
		var rCnt = need[c]
		if 0 != rCnt {
			window[c]++
			if rCnt == window[c] {
				valid++
			}
		}

		// 判断左侧窗口是否收缩
		for right-left >= tLength {
			// 在这里判断是否找到合法子串
			if valid == len(need) {
				return true
			}

			var d = s[left]
			left++

			// 进行窗口内数据的一系列更新
			var lCnt = need[d]
			if 0 != lCnt {
				if lCnt == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
