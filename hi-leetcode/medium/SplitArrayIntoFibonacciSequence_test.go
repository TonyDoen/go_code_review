package medium

import "testing"

func TestSplitIntoFibonacci(t *testing.T) {
	s := "1101111"
	res := SplitIntoFibonacci(s)
	printList(res)
}
