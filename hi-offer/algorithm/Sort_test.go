package algorithm

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	src := "abccdee"
	res := Permutation(src)
	cur, cnt := res.Front(), 0
	for nil != cur {
		print(cur.Value.(string) + " ")
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

func TestQuickSort(t *testing.T) {
	arr := []int{7, 5, 6, 4, 4, 3, 3, 8, 0, 9}
	QuickSort(arr, 0, len(arr)-1)
	for _, v := range arr {
		print(v)
		print(" ")
	}
}

func TestFindKthNumber0(t *testing.T) {
	arr := []int{7, 5, 6, 4, 4, 3, 3, 8, 0, 9}
	res := FindKthNumber0(arr, 5)
	for _, v := range res {
		print(v)
		print(" ")
	}
}

func TestHeapSort(t *testing.T) {
	//arr := []int{7, 5, 6, 4, 4, 3, 3, 8, 0, 9, 10, 11, 11}
	arr := []int{7, 5, 6, 4}
	HeapSort(arr)
	for _, v := range arr {
		print(v)
		print(" ")
	}

}

func TestFindKthNumber1(t *testing.T) {
	arr := []int{7, 5, 6, 4, 4, 3, 3, 8, 0, 9, 10, 11, 11}
	res := FindKthNumber1(arr, 5)
	for _, v := range res {
		print(v)
		print(" ")
	}
}
