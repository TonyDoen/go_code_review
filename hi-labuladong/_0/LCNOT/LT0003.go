package LCNOT

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
 * 1438. 绝对差不超过限制的最长连续子数组
 *
 * 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
 * 如果不存在满足条件的子数组，则返回 0 。
 *
 * 提示：
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 0 <= limit <= 10^9
 *
 *
 * 示例 1：
 * 输入：nums = [8,2,4,7], limit = 4
 * 输出：2
 *
 * 解释：所有子数组如下：
 * [8] 最大绝对差 |8-8| = 0 <= 4.
 * [8,2] 最大绝对差 |8-2| = 6 > 4.
 * [8,2,4] 最大绝对差 |8-2| = 6 > 4.
 * [8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
 * [2] 最大绝对差 |2-2| = 0 <= 4.
 * [2,4] 最大绝对差 |2-4| = 2 <= 4.
 * [2,4,7] 最大绝对差 |2-7| = 5 > 4.
 * [4] 最大绝对差 |4-4| = 0 <= 4.
 * [4,7] 最大绝对差 |4-7| = 3 <= 4.
 * [7] 最大绝对差 |7-7| = 0 <= 4.
 * 因此，满足题意的最长子数组的长度为 2 。
 *
 * 示例 2：
 * 输入：nums = [10,1,2,4,7,2], limit = 5
 * 输出：4
 * 解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
 *
 * 示例 3：
 * 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
 * 输出：3
 *
 *
 *
 * 思路：
 * 滑动窗口
 *
 *
 */
//func LongestSubarray(nums []int, limit int) int {
//	var window, left, right, rs = make(map[int]int), 0, 0, 0
//	for nLength := len(nums); right < nLength; {
//		var ir = nums[right]
//		right++
//
//		// 进行窗口内数据的一系列更新
//		window[ir]++
//
//		// 判断左侧窗口是否收缩
//		for ; window; {
//
//		}
//	}
//}
