package _0

import "math"

/*
 * 滑动窗口模板
 *
 * int left = 0, right = 0;
 * while (right < s.size()) {
 *     // 增大窗口
 *     window.add(s[right]);
 *     right++;
 *
 *     while (window needs shrink) {
 *         // 缩小窗口
 *         window.remove(s[left]);
 *         left++;
 *     }
 * }
 *
 * 时间复杂度是 O(N)
 */

/*
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
*/

/**
 * LeetCode 76 ： Minimum Window Substring
 * 给你一个字符串 s, 一个字符串 t, 在 s 里找到包含t 所有字母的最小子串
 *
 * s = "BANCADBCEFGDAFSDFE"
 * t = "ABC"
 */
func MinWindow(s, t string) string {
	var need, window = make(map[uint8]int), make(map[uint8]int)
	for i, tLength := 0, len(t); i < tLength; i++ {
		need[t[i]]++
	}

	var left, right, valid, start, length, sLength = 0, 0, 0, 0, math.MaxInt64, len(s)
	for right < sLength {
		// c 是将要移入窗口的字符
		var c = s[right]
		// 右移窗口
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
		for valid == len(need) {
			// 这里更新最小覆盖字串
			if right-left < length {
				start = left
				length = right - left
			}

			// d 是将要移出窗口的字符
			var d = s[left]
			// 左移窗口
			left++

			// 进行窗口内数据的一系列更新
			var lCnt = need[c]
			if 0 != lCnt {
				if lCnt == window[d] {
					valid--
				}
				window[d] = window[d] - 1
			}
		}
	}

	if math.MaxInt64 == length {
		return ""
	} else {
		return s[start:length]
	}
}
