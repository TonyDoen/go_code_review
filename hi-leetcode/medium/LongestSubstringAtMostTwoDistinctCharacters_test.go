package medium

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstringTwoDistinct1(t *testing.T) {
	src := "ccaabbbc"
	res := LengthOfLongestSubstringTwoDistinct1(src)
	fmt.Printf("%d\n", res)
}

func TestLengthOfLongestSubstringTwoDistinct2(t *testing.T) {
	src := "ccaabbbc"
	res := LengthOfLongestSubstringTwoDistinct2(src)
	fmt.Printf("%d\n", res)
}
