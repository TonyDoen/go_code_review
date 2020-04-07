package medium

import "container/list"

/**
 * Add Two Numbers
 *
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Example:
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 */
/**
 * 题意：两个数字相加(单链表)
 * 这道并不是什么难题，算法很简单，链表的数据类型也不难，就是建立一个新链表，然后把输入的两个链表从头往后撸，每两个相加，添加一个新节点到新链表后面。为了避免两个输入链表同时为空，我们建立一个 dummy 结点，将两个结点相加生成的新结点按顺序加到 dummy 结点之后，由于 dummy 结点本身不能变，所以用一个指针 cur 来指向新链表的最后一个结点。好，可以开始让两个链表相加了，这道题好就好在最低位在链表的开头，所以可以在遍历链表的同时按从低到高的顺序直接相加。while 循环的条件两个链表中只要有一个不为空行，由于链表可能为空，所以在取当前结点值的时候，先判断一下，若为空则取0，否则取结点值。然后把两个结点值相加，同时还要加上进位 carry。然后更新 carry，直接 sum/10 即可，然后以 sum%10 为值建立一个新结点，连到 cur 后面，然后 cur 移动到下一个结点。之后再更新两个结点，若存在，则指向下一个位置。while 循环退出之后，最高位的进位问题要最后特殊处理一下，若 carry 为1，则再建一个值为1的结点，代码如下：
 */
type lNode struct {
	data int
	next *lNode
}

func AddTwoNumbers1(l1, l2 *lNode) *lNode {
	var res, carry = &lNode{data: -1}, 0
	cur := res

	for nil != l1 || nil != l2 {
		var d1, d2 int
		if nil == l1 {
			d1 = 0
		} else {
			d1 = l1.data
		}
		if nil == l2 {
			d2 = 0
		} else {
			d2 = l2.data
		}
		sum := d1 + d2 + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		cur.next = &lNode{data: sum % 10}
		cur = cur.next
		if nil != l1 {
			l1 = l1.next
		}
		if nil != l2 {
			l2 = l2.next
		}
	}
	if 1 == carry {
		cur.next = &lNode{data: 1}
	}
	return res.next
}

/**
 * Add Two Numbers II
 *
 * You are given two linked lists representing two non-negative numbers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
 *
 * Example:
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
 *
 * 题意：两个数字相加之二
 * 思路：
 * 如果我们给链表翻转一下的话就跟之前的题目一样了，这里我们来看一些不修改链表顺序的方法。由于加法需要从最低位开始运算，而最低位在链表末尾，链表只能从前往后遍历，没法取到前面的元素，那怎么办呢？我们可以利用栈来保存所有的元素，然后利用栈的后进先出的特点就可以从后往前取数字了，我们首先遍历两个链表，将所有数字分别压入两个栈s1和s2中，我们建立一个值为0的res节点，然后开始循环，如果栈不为空，则将栈顶数字加入sum中，然后将res节点值赋为sum%10，然后新建一个进位节点head，赋值为sum/10，如果没有进位，那么就是0，然后我们head后面连上res，将res指向head，这样循环退出后，我们只要看res的值是否为0，为0返回res->next，不为0则返回res即可，
 */
func AddTwoNumber2(l1, l2 *lNode) *lNode {
	st1, st2 := list.New(), list.New()
	for nil != l1 {
		st1.PushFront(l1.data)
		l1 = l1.next
	}
	for nil != l2 {
		st2.PushFront(l2.data)
		l2 = l2.next
	}

	var result *lNode
	for sum := 0; st1.Len()>0 || st2.Len() > 0; {
		if st1.Len() > 0 {
			front := st1.Front()
			sum += front.Value.(int)
			st1.Remove(front)
		}
		if st2.Len() > 0 {
			front := st2.Front()
			sum += front.Value.(int)
			st2.Remove(front)
		}
		if nil == result {
			result = &lNode{data: sum%10, next: nil}
		} else {
			result = &lNode{data: sum%10, next: result}
		}
		sum /= 10
	}
	return result
}
