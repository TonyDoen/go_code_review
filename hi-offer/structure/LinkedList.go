package structure

import "container/list"

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}, next *Node) *Node {
	return &Node{Value: value, Next: next}
}

/**
 * url: https://zhuanlan.zhihu.com/p/56200260
 * 003-从尾到头打印链表
 *
 * 从尾到头打印单链表
 * 输入一个链表，按链表值从尾到头的顺序返回一个 List。
 *
 * 思路：
 * 遍历链表，把元素压入栈中，利用栈后进先出特性，遍历栈中元素，逐个打印
 */
func PrintSingleLinkedList(head *Node) {
	if nil == head {
		return
	}

	stack, cur := list.New(), head
	for nil != cur {
		stack.PushFront(cur.Value)
		cur = cur.Next
	}

	tmp := stack.Front()
	for nil != tmp {
		print(tmp.Value.(int))
		print(" ")
		tmp = tmp.Next()
	}
	println()
}

/**
 * 014-链表中倒数第k个结点
 *
 * 为了能够只遍历一次就能找到倒数第k个节点，可以定义两个指针：
 *（1）第一个指针从链表的头指针开始遍历向前走k-1，第二个指针保持不动；
 *（2）从第k步开始，第二个指针也开始从链表的头指针开始遍历；
 *（3）由于两个指针的距离保持在k-1，当第一个（走在前面的）指针到达链表的尾结点时，第二个指针（走在后面的）指针正好是倒数第k个结点。
 */
func FindKth2Tail(head *Node, k int) *Node {
	if nil == head || k < 0 {
		return nil
	}

	first := head
	for i := k; i > 0; i-- { // 先走k-1步
		if nil == first {
			return nil
		}
		first = first.Next
	}

	second := head
	for {
		second = second.Next
		first = first.Next
		if nil == first {
			break
		}
	}
	return second
}

/**
 * 015-反转链表
 *
 * 输入一个链表，反转链表后，输出新链表的表头。
 *
 * 1、只需遍历一次，遍历时拆开链表，当前元素指向前一个元素
 * 2、使用两个指针，一个指向当前遍历的节点，一个指向新链表的头结点
 */
func ReverseSingleLinkedList(head *Node) *Node {
	var result *Node
	for cur := head; nil != cur; {
		next := cur.Next

		cur.Next = result // 头插
		result = cur

		cur = next
	}
	return result
}

/**
 * 016-合并两个或k个有序链表
 *
 * 合并k个有序链表，不管合并几个，基本还是要两两合并。
 * 1. 合并2个有序链表
 * 2. 依次合并多个
 */
func MergeKthNode(lt *list.List) *Node {
	var result *Node
	cur := lt.Front()
	for nil != cur {
		result = MergeTwoNode(result, cur.Value.(*Node))
		cur = cur.Next()
	}
	return result
}

func MergeTwoNode(n1, n2 *Node) *Node { // 递减
	if nil == n1 {
		return n2
	}
	if nil == n2 {
		return n1
	}
	result := NewNode(-1, nil)
	cur := result
	for nil != n1 && nil != n2 {
		if n1.Value.(int) < n2.Value.(int) {
			cur.Next = n2
			n2 = n2.Next
		} else {
			cur.Next = n1
			n1 = n1.Next
		}
		cur = cur.Next
	}
	if nil != n1 {
		cur.Next = n1
	}
	if nil != n2 {
		cur.Next = n2
	}
	return result.Next
}

type ComplexNode struct {
	Value        interface{}
	Next, Random *ComplexNode
}

/**
 * 025-复杂链表的复制
 *
 * 思路1.
 * 先复制直链，在复制随机指针。 (不用Hash表) 时间复杂度 O(n), 空间复杂度 O(1)。
 *
 * 思路2.
 * 在复制随机指针时，可以用哈希表。          时间复杂度 O(n)；空间复杂度 O(n)。
 *
 * 思路3.
 *    1. 遍历链表，复制每个节点并将该复制后的新节点放至旧节点之后
 *    2. 重新遍历链表，复制旧节点的随机指针给新节点
 *    3. 拆分链表，将链表拆分为原链表和复制后的链表
 *
 * eg:
 *    第一遍： 1 -> 2 -> 3  ====>  1 -> 1' -> 2 -> 2' -> 3 -> 3'
 *    第二遍： 复制随机指针
 *    第三遍： 拆分原来的链表/clone链表
 *
 */
func (cn *ComplexNode) Clone() *ComplexNode {
	if nil == cn {
		return nil
	}
	// 1. 遍历链表，复制每个节点并将该复制后的新节点放至旧节点之后
	cur := cn
	for nil != cur {
		cp := &ComplexNode{Value: cur.Value, Next: cur.Next, Random: cur.Random}
		next := cur.Next
		cur.Next = cp
		cur = next
	}
	// 2. 重新遍历链表，复制旧节点的随机指针给新节点
	cur = cn
	for nil != cur {
		if nil == cur.Random {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 3. 拆分链表，将链表拆分为原链表和复制后的链表
	cur = cn
	cloneHead := cur.Next
	for nil != cur {
		cloneCur := cur.Next
		if nil == cloneCur.Next {
			cloneCur.Next = nil
			cur.Next = nil
		} else {
			next := cloneCur.Next
			cloneCur.Next = cloneCur.Next.Next
			cur.Next = next
		}
		cur = cur.Next
	}
	return cloneHead
}

/**
 * 036-两个链表的第一个公共结点
 *
 * 两个链表的第一个公共结点
 * 输入两个链表，找出它们的第一个公共结点。
 *
 * Y型:
 * 1 -> 2 -> 3
 *            \
 *             6 -> 7 -> 8
 *            /
 *      4 -> 5
 *
 * 分析：
 * 两个链表相交，因为链表元素只有一个指针，故相交后，后面都重合了，即相交只能是Y型，不能是X型相交
 *
 * 思路：
 * 分别遍历两个链表，获取长度。两个指针，一个先走|len1-len2|步，之后在同步遍历，获取第一个相同的元素即可。
 *
 * 思路：
 * 两个链表压入栈中，分别弹出元素，获取最后一个相同的元素即可
 *
 */
func FindFirstCommonNode(h1, h2 *Node) *Node {
	if nil == h1 || nil == h2 {
		return nil
	}

	len1, len2, cur1, cur2 := 0, 0, h1, h2
	for nil != cur1 {
		len1++
		cur1 = cur1.Next
	}
	for nil != cur2 {
		len2++
		cur2 = cur2.Next
	}
	remain := 0
	var fast *Node
	var slow *Node
	if len1 > len2 {
		remain = len1 - len2
		fast = h1
		slow = h2
	} else {
		remain = len2 - len1
		fast = h2
		slow = h1
	}

	for i := 0; i < remain; i++ {
		fast = fast.Next
	}
	for nil != fast && nil != slow {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}

/**
 * 055-链表中环的入口结点
 *
 * 链表中环的入口结点
 * 给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。
 *
 * 思路：
 * 1、快慢指针遍历链表，若相遇，则链表存在环
 * 2、一个指针从相遇的节点出发，一个从链表头部出发，两个指针相遇的位置即为环的入口节点
 *
 * 证明：
 * 1 -> 2 -> 3 -> 4 -> 5 -> 6
 *                |         |
 *                9 <- 8 <- 7
 *
 * 假设 slow  走 1步; fast 走 2步
 * 那么当 slow 和 fast 相遇时，既有  2 * slowDistance = fastDistance
 *
 * 令 链表头到环入口的长度(即上图 1 -> 2 -> 3 -> 4)            是 b
 *    环的长度(即上图 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 4)      是 c
 *    环入口 到 slow 和 fast 相遇点(即上图 4 -> 5 -> 6 -> 7) 是 a
 *    slow 和 fast 相遇点 到 环入口 的长度( -> 8 -> 9 -> 4)  是 c - a
 *
 * 计算
 * 2 * slowDistance    = fastDistance
 * 2 * (b + n * c + a) = (b + m * c + a)
 *                   b = (m-2n-1) * c + (c - a)
 *
 * 即：链表头到环入口的长度                 b
 *    slow 和 fast 相遇点 到 环入口 的长度 (c - a)
 *    相差 (m-2n-1)系数 个 完整的 环的长度  c
 *
 *    所以，思路中第2点得到证明
 *
 */
func FindEntryNodeOfLoop(head *Node) *Node {
	if nil == head || nil == head.Next { // 0,1 个单链表元素构不成环
		return nil
	}
	// 1、快慢指针遍历链表，若相遇，则链表存在环
	var meet *Node
	fast, slow := head, head
	for nil != slow && nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			meet = slow
			break
		}
	}
	if nil == meet { // no loop
		return nil
	}

	// 2、一个指针从相遇的节点出发，一个从链表头部出发，两个指针相遇的位置即为环的入口节点
	cur := head
	for cur != meet {
		cur = cur.Next
		meet = meet.Next
	}
	return cur
}

/**
 * 056-删除链表中重复的结点
 *
 * 删除链表中重复的结点
 * 在一个排序的链表中,存在重复的结点,请删除该链表中重复的结点,重复的结点不保留,返回链表头指针。
 *
 * 例如,链表 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5 处理后为 1 -> 2 -> 5
 *
 * 思路
 * 1、新建前驱节点,next指向pHead,方便处理头节点元素重复的场景
 * 2、双指针,一个pre指针指向前一个不重复的节点,一个cur指向当前遍历的节点,分情况处理
 * 3、当遍历到重复的节点时,pre指针的next指向当前重复节点的最后一个节点的next,即删除重复元素,cur指针后移
 * 4、当遍历到不重复的节点,只需要pre、cur指针同时后移即可
 *
 */
func DeleteDuplicateNode(head *Node) *Node {
	if nil == head {
		return nil
	}
	result := &Node{Value: -1, Next: head}
	pre, cur := result, head
	for nil != cur && nil != cur.Next {
		if cur.Value == cur.Next.Value {
			for nil != cur.Next && cur.Value == cur.Next.Value {
				cur = cur.Next
			}
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return result.Next
}

/**
 * 删除链表中重复的结点, 保留第一个重复结点
 */
func DeleteDuplicateNodeRemainFirstOne(head *Node) *Node {
	if nil == head {
		return nil
	}
	result := NewNode(-1, head)
	pre, cur := result, head
	for nil != pre {
		cur = pre.Next
		if nil != cur && pre.Value == cur.Value {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}
	return result.Next
}
