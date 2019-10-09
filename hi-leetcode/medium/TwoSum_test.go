package medium

import (
	"fmt"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	var nums, target = []int{2, 7, 11, 15}, 9
	res1 := TwoSum1(nums, target)
	fmt.Printf("%d\n", res1)

	res2 := TwoSum2(nums, target)
	fmt.Printf("%d\n", res2)
}
