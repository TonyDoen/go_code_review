package medium

import (
	"fmt"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := findDuplicates(nums)
	tmp := res.Front()
	for nil != tmp {
		fmt.Printf("%d\n", tmp.Value.(int))
		tmp = tmp.Next()
	}
}

func TestFindDuplicates3(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := findDuplicates3(nums)
	tmp := res.Front()
	for nil != tmp {
		fmt.Printf("%d\n", tmp.Value.(int))
		tmp = tmp.Next()
	}
}
