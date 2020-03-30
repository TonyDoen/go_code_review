package medium

import (
	"fmt"
	"testing"
)

func TestLongestPalindromicSubstring1(t *testing.T) {
	src := "fkjlaabbccddeeddccbbaaljkdddddddd"
	res := LongestPalindromicSubstring1(src)
	fmt.Printf("%s\n", res)
}

func TestLongestPalindromicSubstring2(t *testing.T) {
	src := "fkjlaabbccddeeddccbbaaljkdddddddddddddddddddddddddddddddddddddddddddddddddddddd"
	res := LongestPalindromicSubstring2(src)
	fmt.Printf("%s\n", res)
}

func TestLongestPalindromicSubstring3(t *testing.T) {
	src := "fkjlaabbccddeeddccbbaaljkdddddddddddddddddddddddddddddddddddddddddddddddddddddd"
	res := LongestPalindromicSubstring3(src)
	fmt.Printf("%s\n", res)
}
