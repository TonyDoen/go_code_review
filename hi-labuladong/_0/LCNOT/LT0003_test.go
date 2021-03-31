package LCNOT

import (
	"fmt"
	"testing"
)

func TestLongestSubarray1(t *testing.T) {
	var nums, limit = []int{8, 2, 4, 7}, 4

	var rs = LongestSubarray1(nums, limit)
	fmt.Printf("%d\n", rs)
}