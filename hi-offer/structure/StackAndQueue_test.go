package structure

import (
	"strconv"
	"testing"
)

func TestSelfStack(t *testing.T) {
	stack := &SelfStack{}

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for !stack.IsEmpty() {
		println("curMax:" + strconv.Itoa(stack.GetMax().(int)) +
			"; curMin:" + strconv.Itoa(stack.GetMin().(int)) +
			"; curValue:" + strconv.Itoa(stack.Pop().(int)))
	}
}

func TestSelfQueue(t *testing.T) {
	queue := &SelfQueue{}

	for i := 0; i < 10; i++ {
		queue.Offer(i)
	}

	for !queue.IsEmpty() {
		println("curValue:" + strconv.Itoa(queue.Poll().(int)))
	}
}

func TestIsPopOrder(t *testing.T) {
	push, pop := []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}
	res := IsPopOrder(push, pop)
	println(res)
}

func TestReverseSentence0(t *testing.T) {
	src := " student. a am I"
	res := ReverseSentence0(src)
	println(res)
}

func TestReverseSentence1(t *testing.T) {
	src := " student. a am I"
	res := ReverseSentence1(src)
	println(res)
}

func TestMaxInWindow(t *testing.T) {
	arr := []int{2, 3, 4, 2, 6, 2, 5, 1}
	res := MaxInWindow(arr, 3)
	cur := res.Front()
	for nil != cur {
		print(cur.Value.(int))
		print(" ")
		cur = cur.Next()
	}
	println()
}
