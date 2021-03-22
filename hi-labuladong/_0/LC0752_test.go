package _0

import (
	"fmt"
	"go_code_review/hi-leetcode/hard"
	"testing"
)

func TestUnlock(t *testing.T) {
	var deadEnds = hard.NewSet()
	_ = deadEnds.Add("0201", "0101", "0102", "1212", "2002")

	var target = "0202"

	var rs = Unlock(deadEnds, target)
	fmt.Printf("%d\n", rs)
}
