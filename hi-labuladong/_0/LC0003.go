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
 * LeetCode 第3题 Longest Substring Without Repeating Characters
 *
 * 给定一个字符串，请你找出其中不含重复字符的最长子串的长度
 *
 * 例子1:
 * input: "abcabcbb"
 * output: 3
 * explain: 因为无重复字符的最长子串 "abc"，所以长度为3
 *
 * 例子2:
 * input: "bbbbb"
 * output: 1
 * explain: 因为无重复字符的最长子串 "b"，所以长度为1
 *
 * 例子3:
 * input: "pwwkew"
 * output: 3
 * explain: 因为无重复字符的最长子串 "wke"，所以长度为3
 *
 */
func LengthOfLongestSubstring(s string) int {
	var window, rs = make(map[uint8]int), 0

	for sLength, left, right := len(s), 0, 0; right < sLength; {
		var c = s[right]
		right++

		// 进行窗口内数据的一系列更新
		window[c]++

		// 判断左侧窗口是否收缩
		for window[c] > 1 {
			var d = s[left]
			left++

			// 进行窗口内数据的一系列更新
			window[d]--
		}
		rs = max(rs, right-left)
	}
	return rs
}

func max(args ...int) int {
	var rs = math.MinInt64
	for _, v := range args {
		if rs < v {
			rs = v
		}
	}
	return rs
}
