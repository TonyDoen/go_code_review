package LCNOT

import (
	"fmt"
	"testing"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	var intervals = [][]int{{1, 4}, {3, 6}, {2, 8}}
	var rs = RemoveCoveredIntervals(intervals)
	fmt.Printf("%d\n", rs)
}
