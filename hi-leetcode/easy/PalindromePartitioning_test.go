package easy

import (
	"container/list"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	res := isPalindrome("cabgbac")
	println(res)
}

// PalindromePartitioning
func TestPalindromePartitioning(t *testing.T) {
	res := PalindromePartitioning("cabgbac")
	tmp := res.Front()

	print("[\n")
	for nil != tmp {
		tmpL := tmp.Value.(*list.List)
		cur := tmpL.Front()
		print("[")
		for nil != cur {
			print(cur.Value.(string))
			cur = cur.Next()
			if nil != cur {
				print(", ")
			}
		}
		print("]")
		tmp = tmp.Next()
		if nil != tmp {
			print(",\n")
		}
	}
	println("\n]")
}

// PalindromePartitioning2
func TestMinCut(t *testing.T) {
	res := MinCut("aab")
	println(res)
}

func TestMinCut0(t *testing.T) {
	res := MinCut0("aab")
	println(res)
}
