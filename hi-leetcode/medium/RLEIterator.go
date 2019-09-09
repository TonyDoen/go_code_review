package medium

import "container/list"

type RLEIteratorOne struct {
	seq list.List
}

func NewRLEIteratorOne(arr []int) *RLEIteratorOne {
	length := len(arr)
	var seq list.List
	for i := 0; i < length; i += 2 {
		if 0 != arr[i] {
			seq.PushBack(NewPair(arr[i+1], arr[i]))
		}
	}
	return &RLEIteratorOne{seq: seq}
}

func (ro *RLEIteratorOne) next(n int) int { // error
	if nil == ro {
		return -1
	}
	//tmp := ro.seq.Front()
	//for {
	//	p := tmp.Value.(*Pair)
	//	//if 0 == p.v {
	//	//	continue
	//	//}
	//	if n <= p.v.(int) {
	//		p.v = p.v.(int) - n
	//		return p.k.(int)
	//	}
	//	n -= p.v.(int)
	//
	//	tmp := tmp.Next()
	//	if nil == tmp || nil == tmp.Value {
	//		break
	//	}
	//}
	return -1
}
