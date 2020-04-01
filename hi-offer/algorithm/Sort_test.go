package algorithm

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	src := "abccdee"
	res := Permutation(src)
	cur, cnt := res.Front(), 0
	for nil != cur {
		print(cur.Value.(string)+" ")
		cur = cur.Next()
		cnt++
	}
	println(cnt)
}

func TestInversePair(t *testing.T) {
	arr := []int{7, 5, 6, 4}
	res := InversePair(arr)
	println(res)
}