package medium

import "testing"

func TestFindReplaceString(t *testing.T) {
	src, indexes, sources, targets := "abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}
	result := FindReplaceString(src, indexes, sources, targets)
	println(result)
}
