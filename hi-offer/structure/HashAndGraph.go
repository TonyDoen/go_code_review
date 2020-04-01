package structure

import (
	"strconv"
	"strings"
)

/**
 * LinkedHashMap
 * 实现按输入顺序排序的HashMap
 */
type LinkedHashMap struct {
	mp         map[interface{}]interface{}
	head, tail *Node
}

func (l *LinkedHashMap) Get(key interface{}) interface{} {
	if nil == l.mp {
		return nil
	}
	return l.mp[key]
}

func (l *LinkedHashMap) Put(key, value interface{}) {
	if nil == key || nil == value {
		return
	}
	if nil == l.mp {
		l.mp = make(map[interface{}]interface{})
		l.mp[key] = value
		l.head = NewNode(key, nil)
		l.tail = l.head
		return
	}
	oldValue := l.mp[key]
	if oldValue == value {
		return
	}
	l.mp[key] = value
	if nil == oldValue {
		l.tail.Next = NewNode(key, nil)
		l.tail = l.tail.Next
	}
}

func (l *LinkedHashMap) Remove(key interface{}) {
	if nil == key || nil == l.mp {
		return
	}
	oldValue := l.mp[key]
	if nil == oldValue {
		return
	}
	delete(l.mp, key)
	cur, pre := l.head, l.head
	for nil != cur {
		if key == cur.Value {
			if l.head == cur {
				l.head = l.head.Next
				break
			}
			pre.Next = cur.Next
			if l.tail == cur {
				l.tail = pre
			}
			break
		}

		pre = cur
		cur = cur.Next
	}
	if nil == l.head && nil != l.tail {
		l.tail = nil
	}
}

func (l *LinkedHashMap) Next() *Node {
	if nil == l.head {
		return nil
	}
	return l.head
}

/**
 * HashTable
 * 034-第一个只出现一次的字符
 *
 * 题目：
 * 第一个只出现一次的字符
 *
 * 题目描述：
 * 在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,
 * 并返回它的位置, 如果没有则返回 -1（需要区分大小写）.
 *
 * 思路：
 * 1、遍历字符串，Hash存储字符串每个字符出现的次数
 * 2、顺序遍历上面存储的结果，如果该字符出现次数为1次，则找到该字符及其位置
 */
func FirstNotRepeatingChar(src string) int {
	length := len(src)
	if length < 1 {
		return -1
	}

	mp := &LinkedHashMap{}
	for i := 0; i < length; i++ {
		c := src[i]
		idxCnt := mp.Get(c)
		if nil == idxCnt {
			mp.Put(c, "1|"+strconv.Itoa(i))
		} else {
			piece := strings.Split(idxCnt.(string), "|")
			cnt, _ := strconv.Atoi(piece[0])
			mp.Put(c, strconv.Itoa(cnt+1)+"|"+piece[1])
		}
	}

	cur := mp.Next()
	for nil != cur {
		v := mp.Get(cur.Value).(string)
		piece := strings.Split(v, "|")
		cnt, _ := strconv.Atoi(piece[0])
		if 1 == cnt {
			idx, _ := strconv.Atoi(piece[1])
			return idx
		}
		cur = cur.Next
	}
	return -1
}

/**
 * 065-矩阵中的路径(回溯法)
 *
 * 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
 * 路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左，向右，向上，向下移动一个格子。
 * 如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
 * a b c e
 * s f c s
 * a d e e
 * 矩阵中包含一条字符串"bcced“的路径，但是矩阵中不包含”abcb"路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入该格子。
 *
 * 思路：
 * 1. 首先确定利用回溯法进行求解。确立解空间-子集树空间。
 * 2. 利用深度优先搜索。
 * 3. 确立剪枝条件。
 *    1>搜索越出矩阵范围，停止向下搜索，剪枝。
 *    2>不是目标字符，停止向下搜索，剪枝。
 *    3>已经被搜索比较过，停止向下所搜，剪枝。
 */
func HasPath(matrix [][]rune, target []rune) bool {
	if nil == matrix || nil == target || len(matrix) < 1 || len(matrix[0]) < 1 || len(target) < 1 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	visited := make([]bool, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if hasPath(0, row, col, i, j, visited, target, matrix) {
				return true
			}
		}
	}
	return false
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 右左下上
func hasPath(start, row, col, i, j int, visited []bool, target []rune, matrix [][]rune) bool {
	idx := i*col + j // 先根据i和j计算匹配的第一个元素转为一维数组的位置
	// 递归终止条件
	if i < 0 || j < 0 || i >= row || j >= col || matrix[i][j] != target[start] || true == visited[idx] {
		return false
	}
	if len(target)-1 == start { // 若k已经到达str末尾了，说明之前的都已经匹配成功了，直接返回true即可
		return true
	}

	visited[idx] = true // 要走的第一个位置置为true，表示已经走过了
	for _, dir := range dirs { // 回溯，递归寻找，每次找到了就给k加一，找不到，还原
		if hasPath(start+1, row, col, i+dir[0], j+dir[1], visited, target, matrix) {
			return true
		}
	}
	visited[idx] = false // 走到这，说明这一条路不通，还原，再试其他的路径; 回溯需要还原
	return false
}

/**
 * 066-机器人的运动范围
 *
 * 地上有一个m行和n列的方格。一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向移动一格，但是不能进入行坐标和列坐标的数位之和大于k的格子。
 * 例如，当k为18时，机器人能够进入方格（35,37），因为3+5+3+7 = 18。但是，它不能进入方格（35,38），因为3+5+3+8 = 19。
 * 请问该机器人能够达到多少个格子？
 *
 * 1. 依然利用回溯法的子集树解空间思想。 确定解空间-子集树。
 * 2. 利用深度优先搜索。
 * 3. 确定剪枝条件。
 *    1>搜索到矩阵边界外，剪枝。
 *    2>该位置已经被搜索过了，剪枝。
 *    3>和大于K的，剪枝。
 * 需注意 的一点是，这里搜索过的位置不需要回溯了，也就是说不在搜索这里。
 */
func CountMove(threshold, row, col int) int {
	visited := make([]bool, row*col)
	return count(visited, threshold, 0, 0, row, col)
}
func count(visited []bool, threshold, i, j, row, col int) int {
	idx := i*col + j // 先根据i和j计算匹配的第一个元素转为一维数组的位置
	if i < 0 || j < 0 || i > row || j > col || sumInt(i)+sumInt(j) > threshold || true == visited[idx] {
		return 0
	}
	visited[idx] = true
	result := 1
	for _, dir := range dirs { // 回溯，递归寻找，
		result += count(visited, threshold, i+dir[0], j+dir[1], row, col)
	}
	return result
}
func sumInt(i int) int {
	result := 0
	for i > 0 {
		result += i % 10
		i = i / 10
	}
	return result
}
