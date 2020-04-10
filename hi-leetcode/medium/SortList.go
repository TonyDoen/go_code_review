package medium

/**
 * Sort a linked list in O(n log n) time using constant space complexity.
 *
 * Example 1:
 * Input: 4->2->1->3
 * Output: 1->2->3->4
 *
 * Example 2:
 * Input: -1->5->3->4->0
 * Output: -1->0->3->4->5
 *
 *
 * 题意：给链表排序
 *
 * 思路：
 * 由于时间复杂度要求o(n Log n),空间复杂度要求固定，则需要使用归并排序
 * 1、分解链表时，获取中间节点后，可以把中间节点的next指向null，便于后续获取中间节点及归并处理
 * 2、合并两个有序链表时，可定义要返回的新链表的头节点的pre节点，避免确定新链表头节点时的比较操作
 *
 * 时间复杂度要求O(n Log n)
 */
// 归并排序
func SortLinkedList0(head *lNode) *lNode {
	if nil == head || nil == head.next {
		return head
	}
	// 快慢指针
	fast, slow := head, head
	for nil != fast.next && nil != fast.next.next {
		fast = fast.next.next
		slow = slow.next
	}
	mid := slow
	var another *lNode
	if nil == mid {
		another = nil
	} else {
		another = mid.next
		mid.next = nil // 变原链表为两个独立的链表，很巧妙
	}
	return mergeSortLinkedList(SortLinkedList0(head), SortLinkedList0(another))
}
func mergeSortLinkedList(first, second *lNode) *lNode { // 合并两个有序链表为一个链表
	if nil == first && nil == second {
		return nil
	}
	if nil == first {
		return second
	}
	if nil == second {
		return first
	}
	pre := &lNode{-1, nil}
	curNode, curF, curS := pre, first, second
	for nil != curF && nil != curS {
		if curF.data <= curS.data {
			curNode.next = curF
			curF = curF.next
		} else {
			curNode.next = curS
			curS = curS.next
		}
		curNode = curNode.next
	}
	if nil != curF {
		curNode.next = curF
	}
	if nil != curS {
		curNode.next = curS
	}
	return pre.next
}

/**
 * 快速排序
 *
 * 时间复杂度为 O(n Log n);
 *
 * 思路：
 * 1、以头节点为base，依次遍历，与base比较，比base小的放前面；返回base指针
 * 2、递归调用上述过程
 */
// 快速排序
func SortLinkedList1(head *lNode) *lNode {
	if nil == head || nil == head.next {
		return head
	}
	quickSortLinkedList(head, nil)
	return head
}
func quickSortLinkedList(begin, end *lNode) {
	if begin == end || begin.next == end {
		return
	}
	// partition
	baseData, base, cur := begin.data, begin, begin.next
	for cur != end {
		if cur.data < baseData {
			base = base.next
			// swap(base.data, cur.data)
			tmp := base.data
			base.data = cur.data
			cur.data = tmp
		}
		cur = cur.next
	}
	// swap(base.data, begin.data)
	tmp := base.data
	base.data = begin.data
	begin.data = tmp

	//
	pivot := base
	quickSortLinkedList(begin, pivot)
	quickSortLinkedList(pivot.next, end)
}
