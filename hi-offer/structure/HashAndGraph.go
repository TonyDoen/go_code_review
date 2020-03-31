package structure

import (
	"strconv"
	"strings"
)

/**
 * LinkedHashMap
 * 实现按输入顺序排序的HashMap
 *
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
