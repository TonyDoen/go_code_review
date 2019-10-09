package medium

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstringKDistinct2(t *testing.T) {
	var s, k = "eceba", 2

	res := LengthOfLongestSubstringKDistinct2(s, k)
	fmt.Printf("%d\n", res)
}
