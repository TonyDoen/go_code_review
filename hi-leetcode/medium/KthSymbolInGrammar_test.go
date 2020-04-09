package medium

import "testing"

func TestKthGrammar0(t *testing.T) {
	n, k := 4, 5
	res := KthGrammar0(n, k)
	println(res)

	res = KthGrammar0Up(n, k)
	println(res)
}

func TestKthGrammar2(t *testing.T) {
	n, k := 4, 5
	res := KthGrammar2(n, k)
	println(res)

	res = KthGrammar2Up(n, k)
	println(res)
}
