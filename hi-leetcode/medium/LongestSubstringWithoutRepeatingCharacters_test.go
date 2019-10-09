package medium

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring1(t *testing.T) {
	src := "abcabcbb"
	res := LengthOfLongestSubstring1(src)
	fmt.Printf("%d\n", res)

	res = LengthOfLongestSubstring2(src)
	fmt.Printf("%d\n", res)
}
