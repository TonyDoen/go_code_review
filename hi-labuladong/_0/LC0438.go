package _0

import "container/list"

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

  时间复杂度是 O(N)
*/

/*
 * LeetCode 438 Find All Anagrams in a String
 *
 * 给定一个字符串s, 一个非空字符串p, 找到s中所有是p的字母异位词的子串, 返回这些字串的起始索引。
 * 字符串只包含小写英文字母，并且字符串s和p的长度都不超过20100
 *
 * 说明
 * <1>字母异位词：字母相同，但是排序不一定相同的字符串
 * <2>不考虑答案输出顺序
 *
 * 例子1：
 * input: s = "cbaebabacd"; p = "abc"
 * output: [0, 6]
 * explain: 索引位置0的子串 "cba", 是 "abc" 字母异位词
 *          索引位置6的子串 "bac", 是 "abc" 字母异位词
 *
 * 相当于，输入s, t; 找到 s 中所有 t 的排列，返回他们的起始索引
 *
 */
func FindAnagrams(s, t string) *list.List {
	var need, window, tLength = make(map[uint8]int), make(map[uint8]int), len(t)
	for i := 0; i < tLength; i++ {
		need[t[i]]++
	}

	var left, right, valid, rs = 0, 0, 0, list.New()
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
			// 当窗口符合条件，把起始索引加入res
			if valid == len(need) {
				rs.PushBack(left)
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
	return rs
}
