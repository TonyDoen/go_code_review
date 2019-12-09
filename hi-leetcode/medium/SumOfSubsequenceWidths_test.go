package medium

import (
	"fmt"
	"testing"
)

func TestSumSubseqWidths1(t *testing.T) {
	var nums = []int{2, 1, 3}
	res1 := SumSubseqWidths1(nums)
	fmt.Printf("%d\n", res1)
}

func TestSumSubseqWidths2(t *testing.T)  {
	var nums = []int{2, 1, 3}
	res1 := SumSubseqWidths2(nums)
	fmt.Printf("%d\n", res1)
}