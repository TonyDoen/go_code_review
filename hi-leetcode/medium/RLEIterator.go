package medium

import "container/list"

type RLEIterator1 struct {
	seq *list.List
}

func NewRLEIterator1(arr []int) *RLEIterator1 {
	length, seq := len(arr), list.New()
	for i := 0; i < length; i += 2 {
		if 0 != arr[i] {
			seq.PushBack(NewPair(arr[i+1], arr[i])) // 前一个数字表示后面的一个数字重复出现的次数
		}
	}
	return &RLEIterator1{seq: seq}
}

func (rr *RLEIterator1) next(n int) int {
	if nil == rr || nil == rr.seq {
		return -1
	}
	tmp := rr.seq.Front()
	for nil != tmp {
		p := tmp.Value.(*Pair)
		k, v := p.k.(int), p.v.(int)

		tmp = tmp.Next()
		if 0 == v {
			continue
		}
		if v >= n {
			p.v = v - n
			return k
		}
		n -= v
		p.v = 0
	}
	return -1
}
