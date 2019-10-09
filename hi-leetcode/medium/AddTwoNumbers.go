package medium

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
		cur.next = &lNode{data: sum%10}
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
