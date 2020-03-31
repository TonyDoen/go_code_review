package structure

import (
	"container/list"
	"math"
)

/**
 * 005-用两个栈实现队列
 */
type QueueByStack struct {
	_1 *SelfStack
	_2 *SelfStack
}

func (q *QueueByStack) Offer(value interface{}) {
	q._1.Push(value)
}

func (q *QueueByStack) Poll() interface{} {
	if q._1.IsEmpty() && q._2.IsEmpty() {
		return nil
	}
	if !q._2.IsEmpty() {
		return q._2.Pop()
	}
	for !q._1.IsEmpty() {
		q._2.Push(q._1.Pop())
	}
	return q._2.Pop()
}

/**
 * 006-自己实现队列
 */
type SelfQueue struct {
	head *Node
	tail *Node
}

func (s *SelfQueue) Offer(value interface{}) {
	if nil == s.head {
		s.head = NewNode(value, nil)
		s.tail = s.head
		return
	}
	s.tail.Next = NewNode(value, nil)
	s.tail = s.tail.Next
}

func (s *SelfQueue) Poll() interface{} {
	if nil == s.head {
		return nil
	}
	result := s.head.Value
	s.head = s.head.Next
	if nil == s.head {
		s.tail = nil
	}
	return result
}

func (s *SelfQueue) IsEmpty() bool {
	return nil == s.head
}

/**
 * 020-包含min/max函数的栈
 */
type SelfStack struct {
	head *Node
	max  *Node
	min  *Node
}

func (s *SelfStack) Push(value interface{}) {
	if nil == s.head {
		s.head = NewNode(value, nil)
		s.max = s.head
		s.min = s.head
		return
	}
	s.head = NewNode(value, s.head)
	// max
	if s.max.Value.(int) < value.(int) { // TODO：优化
		s.max = NewNode(value, s.max)
	} else {
		s.max = NewNode(s.max.Value, s.max)
	}
	// min
	if s.min.Value.(int) > value.(int) { // TODO：优化
		s.min = NewNode(value, s.min)
	} else {
		s.min = NewNode(s.min.Value, s.min)
	}
}

func (s *SelfStack) Pop() interface{} {
	if nil == s.head {
		return nil
	}
	result := s.head.Value
	s.head = s.head.Next
	s.max = s.max.Next
	s.min = s.min.Next
	return result
}

func (s *SelfStack) GetMax() interface{} {
	if nil == s.max {
		return nil
	}
	return s.max.Value
}

func (s *SelfStack) GetMin() interface{} {
	if nil == s.min {
		return nil
	}
	return s.min.Value
}

func (s *SelfStack) Peek() interface{} {
	if nil == s.head {
		return nil
	}
	return s.head.Value
}

func (s *SelfStack) IsEmpty() bool {
	return nil == s.head
}

/**
 * 021-栈的压入、弹出序列
 *
 * 栈的压入、弹出序列
 *
 * 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
 * 假设压入栈的所有数字均不相等。
 * 例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，
 * 但4,3,5,1,2就不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）
 *
 *
 * 思路1：
 * 基于元素入栈序列，模拟出栈过程，最后还有元素残留，即为不是正确出栈序列
 * 1、push 元素依次入栈，必须先入栈元素，否则最后一个元素无法正确处理
 * 2、循环--如果栈顶元素是要出栈的元素，则弹出，出栈元素索引+1
 * 3、最后返回，栈是否为空 即可
 *
 */
func IsPopOrder(push, pop []int) bool {
	if nil == push || nil == pop || len(push) != len(pop) || len(push) < 1 {
		return false
	}
	stack, length := &SelfStack{}, len(push)
	for i, j := 0, 0; i < length; i++ {
		stack.Push(push[i])
		for !stack.IsEmpty() && stack.Peek().(int) == pop[j] { // 栈顶元素=要弹出的元素，则栈中元素出栈
			stack.Pop()
			j++
		}
	}
	return stack.IsEmpty()
}

/**
 * 044-翻转单词顺序列(栈)
 * reverseSentence
 * 最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。
 * 同事Cat对Fish写的内容颇感兴趣，有一天他向Fish借来翻看，但却读不懂它的意思。
 * 例如，“student. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，
 * 正确的句子应该是“I am a student.”。
 *
 * 思路1：
 * 1、stack
 *
 */
func ReverseSentence0(src string) string {
	length := len(src)
	if length < 1 {
		return ""
	}
	stack := list.New()
	pre := 0
	for i := 0; i < length; i++ {
		if ' ' == src[i] {
			stack.PushFront(src[pre:i])
			pre = i
		}
	}
	if pre < length {
		stack.PushFront(src[pre:length])
	}
	result := ""
	for stack.Len() > 0 {
		front := stack.Front()
		result += front.Value.(string)
		stack.Remove(front)
	}
	return result
}

/**
 * 思路2：
 * 1、先把每个单词都进行反转
 * 2、再把字符串全部反转
 */
func ReverseSentence1(src string) string {
	length := len(src)
	if length < 1 {
		return ""
	}
	arr := []byte(src)
	pre := 0
	for i := 0; i < length; i++ {
		if ' ' == src[i] {
			reverse(arr, pre, i-1)
			pre = i
		}
	}
	reverse(arr, 0, length-1)
	return string(arr)
}

func reverse(src []byte, start, end int) {
	if end-start < 1 {
		return
	}
	for start < end {
		cur := src[start]
		src[start] = src[end]
		src[end] = cur
		start++
		end--
	}
}

/**
 * 064-滑动窗口的最大值(双端队列)
 *
 * 给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。
 * 例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，
 * 那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；
 * 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
 * {[2,3,4],2,6,2,5,1}，
 * {2,[3,4,2],6,2,5,1}，
 * {2,3,[4,2,6],2,5,1}，
 * {2,3,4,[2,6,2],5,1}，
 * {2,3,4,2,[6,2,5],1}，
 * {2,3,4,2,6,[2,5,1]}。
 */
/**
 * 思路：
 * 1、遍历数组，每次取滑动窗口的所有值，遍历窗口中的元素，取出最大值
 */
func MaxInWindow(arr []int, size int) *list.List {
	if nil == arr || size < 1 {
		return nil
	}

	result, length := list.New(), len(arr)
	if length <= size {
		max := math.MinInt64
		for _, v := range arr {
			if v > max {
				max = v
			}
		}
		result.PushBack(max)
		return result
	}

	step := length - size
	for i := 0; i <= step; i++ {
		curMax := arr[i]
		for j := i + 1; j < i+size; j++ {
			if curMax < arr[j] {
				curMax = arr[j]
			}
		}
		result.PushBack(curMax)
	}
	return result
}
