package medium

import "container/list"

/**
 * url: https://www.cnblogs.com/grandyang/p/10415773.html
 * url: https://leetcode.com/problems/keys-and-rooms/discuss/133944/Java-8-lines
 *
 * There are N rooms and you start in room 0.  Each room has a distinct number in 0, 1, 2, ..., N-1, and each room may have some keys to access the next room.
 * Formally, each room i has a list of keys rooms[i], and each key rooms[i][j] is an integer in [0, 1, ..., N-1] where N = rooms.length.  A key rooms[i][j] = v opens the room with number v.
 * Initially, all the rooms start locked (except for room 0).
 * You can walk back and forth between rooms freely.
 * Return true if and only if you can enter every room.
 *
 * Example 1:
 * Input: [[1],[2],[3],[]]
 * Output: true
 *
 * Explanation:
 * We start in room 0, and pick up key 1.
 * We then go to room 1, and pick up key 2.
 * We then go to room 2, and pick up key 3.
 * We then go to room 3.  Since we were able to go to every room, we return true.
 *
 * Example 2:
 * Input: [[1,3],[3,0,1],[2],[0]]
 * Output: false
 *
 * Explanation: We can't enter the room with number 2.
 *
 * Note:
 * 1 <= rooms.length <= 1000
 * 0 <= rooms[i].length <= 1000
 * The number of keys in all rooms combined is at most 3000.
 *
 *
 * 题意：钥匙与房间
 * 这道题给了我们一些房间，房间里有一些钥匙，用钥匙可以打开对应的房间，说是起始时在房间0，问我们最终是否可以打开所有的房间。
 * 思路：
 * 这是一道典型的有向图的遍历的题，邻接链表都已经帮我们建立好了，我们直接遍历就好了，这里先用BFS来遍历。使用一个HashSet来记录访问过的房间，先把0放进去，然后使用queue来辅助遍历，同样将0放入。之后进行典型的BFS遍历，取出队首的房间，然后遍历其中的所有钥匙，若该钥匙对应的房间已经遍历过了，直接跳过，否则就将钥匙加入HashSet。此时我们看若HashSet中的钥匙数已经等于房间总数了，直接返回true，因为这表示所有房间已经访问过了，否则就将钥匙加入队列继续遍历。最后遍历结束后，就看HashSet中的钥匙数是否和房间总数相等即可，
 *
 */
// DFS
func CanVisitAllRooms0(rooms [][]int) bool {
	if nil == rooms {
		return false
	}
	visited := NewSet()
	visitRoom0(visited, 0, rooms)
	return visited.Size() == len(rooms)
}
func visitRoom0(visited *Set, start int, rooms [][]int) {
	_ = visited.Add(start)
	for _, roomKey := range rooms[start] {
		if visited.Contains(roomKey) {
			continue
		}
		visitRoom0(visited, roomKey, rooms)
	}
}

// BFS 非递归写法
func CanVisitAllRooms1(rooms [][]int) bool {
	if nil == rooms {
		return false
	}

	visited, queue, length := NewSet(), list.New(), len(rooms)
	// 题意可知, 起始时在房间0
	_ = visited.Add(0)
	queue.PushBack(0)

	for queue.Len() > 0 {
		front := queue.Front()
		num := front.Value.(int)
		queue.Remove(front)

		for _, roomKey := range rooms[num] {
			if visited.Contains(roomKey) {
				continue
			}
			if visited.Size() == length {
				return true
			}
			_ = visited.Add(roomKey)
			queue.PushBack(roomKey)
		}
	}
	return visited.Size() == length
}
