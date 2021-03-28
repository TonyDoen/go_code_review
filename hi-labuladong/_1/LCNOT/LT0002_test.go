package LCNOT

import (
	"fmt"
	"testing"
)

// 区间覆盖问题
func TestRemoveCoveredIntervals(t *testing.T) {
	var intervals = [][]int{{1, 4}, {3, 6}, {2, 8}}
	var rs = RemoveCoveredIntervals(intervals)
	fmt.Printf("%d\n", rs)
}

// 区间合并问题
func TestMergeCoveredIntervals(t *testing.T) {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	var rs = MergeCoveredIntervals(intervals)

	for one := rs.Front(); nil != one; one = one.Next() {
		var arr = one.Value.([]int)
		fmt.Printf("[%d, %d]\n", arr[0], arr[1])
	}
}

// 区间交集问题
func TestIntervalIntersection(t *testing.T) {
	var src0, src1 = [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	var rs = IntervalIntersection(src0, src1)
	for one := rs.Front(); nil != one; one = one.Next() {
		var arr = one.Value.([]int)
		fmt.Printf("[%d, %d]\n", arr[0], arr[1])
	}
}
