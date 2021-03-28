package _0

import (
	"bytes"
	"container/list"
	"go_code_review/hi-leetcode/hard"
)

/**
 * 解开密码锁的最少次数
 *
 *
 * LeetCode 第 752 题: 打开转盘锁
 * 有个4个圆盘的转盘锁，每个转盘有10个数字：0->9;每个拨轮可以自由转动，每次旋转只能转一个拨轮的一位数字，初始'0000'，
 * 列表包含 deadends 死亡数字，一旦拨轮和列表中任何一个元素相同，锁被锁死，无法转动。
 * 字符窜target 代表可以解锁的数字，给出需要转动的最小次数，如果不能解锁，返回-1
 *
 * 例子1
 * 输入： deadends=["0201","0101","0102","1212","2002"], target="0202"
 * 输出： 6
 *
 * 思路：
 * 第一步,我们不管所有的限制条件,不管 deadends 和 target 的限制,就思考一个问题:如果让你设计一个算法,穷举所有可能的密码组合,你怎么做?
 * 穷举呗,再简单一点,如果你只转一下锁,有几种可能?总共有 4 个位置, 每个位置可以向上转,也可以向下转,也就是有 8 种可能对吧。
 * 比如说从 "0000" 开始,转一次,可以穷举出 "0900"... 共 8 种密码。然后,再以这 8 种密码作为基础,对每个密码再转 "1000", "9000", "0100", 一下,穷举出所有可能...
 *
 * 仔细想想,这就可以抽象成一幅图,每个节点有 8 个相邻的节点,又让你求最短距离,
 * 这不就是典型的 BFS 嘛,框架就可以派上用场了,先写出一个 「简陋」的 BFS 框架代码再说
 *
 *
 *
 * 这段 BFS 代码已经能够穷举所有可能的密码组合了,但是显然不能完成题目,有如下问题需要解决:
 * 1、会走回头路。
 * 比如说我们从 "1000" 时,还会拨出一个 "0000"
 * "0000" 拨到 "1000" ,但是等从队列拿出 ,这样的话会产生死循环。
 *
 * 2、没有终止条件,
 * 按照题目要求,我们找到 target 就应该结束并返回拨动的次数。
 *
 * 3、没有对 deadends 的处理,
 * 按道理这些「死亡密码」是不能出现的,也就是说你遇到这些密码的时候需要跳过。
 *
 * 如果你能够看懂上面那段代码,真得给你鼓掌,只要按照 BFS 框架在对应的位置稍作修改即可修复这些问题:
 */
func Unlock(deadEnds *hard.Set, target string) int {
	var visited = hard.NewSet()
	var q = list.New()
	_ = visited.Add("0000")
	q.PushBack("0000")

	var step = 0
	for q.Len() > 0 {
		var sz = q.Len()
		// expand from this node
		for i := 0; i < sz; i++ {
			var front = q.Front()
			var cur = front.Value.(string)
			q.Remove(front)

			if deadEnds.Contains(cur) {
				continue
			}
			if cur == target {
				return step
			}

			// add [cur]'s neighbors
			for j := 0; j < 4; j++ {
				var up = plus1(cur, j)
				if !visited.Contains(up) {
					_ = visited.Add(up)
					q.PushBack(up)
				}

				var down = minus1(cur, j)
				if !visited.Contains(down) {
					_ = visited.Add(down)
					q.PushBack(down)
				}
			}
		}
		step++
	}

	return -1
}

func plus1(s string, j int) string {
	return doPlusPN1(s, j, true)
}

func minus1(s string, j int) string {
	return doPlusPN1(s, j, false)
}

func doPlusPN1(s string, j int, tf bool) string {
	var buff = new(bytes.Buffer)
	var length = len(s)
	for i := 0; i < length; i++ {
		if i == j {
			if tf {
				if '9' == s[i] {
					buff.WriteByte('0')
				} else {
					buff.WriteByte(s[i] + 1)
				}
			} else {
				if '0' == s[i] {
					buff.WriteByte('9')
				} else {
					buff.WriteByte(s[i] - 1)
				}
			}
		} else {
			buff.WriteByte(s[i])
		}
	}
	return buff.String()
}
