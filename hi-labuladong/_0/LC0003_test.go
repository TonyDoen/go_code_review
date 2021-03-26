package _0

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var s = "pwwkew"
	var rs = LengthOfLongestSubstring(s)
	fmt.Printf("%d\n", rs)
}
