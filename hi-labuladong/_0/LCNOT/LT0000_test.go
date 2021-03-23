package LCNOT

import (
	"fmt"
	"testing"
)

func TestBinarySearch0(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 3}
	var target = 2
	var rs = BinarySearch0(nums, target)
	fmt.Printf("%d\n", rs)
}

func TestLeftBinarySearch0(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 3}
	var target = 2
	var rs = LeftBinarySearch0(nums, target)
	fmt.Printf("%d\n", rs)
}

func TestRightBinarySearch0(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 3}
	var target = 2
	var rs = RightBinarySearch0(nums, target)
	fmt.Printf("%d\n", rs)
}

func TestNormalBinarySearch1(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 3}
	var target = 2
	var rs = NormalBinarySearch1(nums, target)
	fmt.Printf("%d\n", rs)

}

func TestLeftBinarySearch1(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 3}
	var target = 2
	var rs = LeftBinarySearch1(nums, target)
	fmt.Printf("%d\n", rs)
}

func TestRightBinarySearch1(t *testing.T) {
	var nums = []int{1, 2, 2, 2, 3}
	var target = 2
	var rs = RightBinarySearch1(nums, target)
	fmt.Printf("%d\n", rs)
}

