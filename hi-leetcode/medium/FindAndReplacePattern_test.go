package medium

import (
	"container/list"
	"testing"
)

func TestFindAndReplacePattern0(t *testing.T) {
	words, pattern := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"
	res := FindAndReplacePattern0(words, pattern)
	printStringList(res)
}

func TestFindAndReplacePattern1(t *testing.T) {
	words, pattern := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"
	res := FindAndReplacePattern1(words, pattern)
	printStringList(res)
}

func TestFindAndReplacePattern2(t *testing.T) {
	words, pattern := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"
	res := FindAndReplacePattern2(words, pattern)
	printStringList(res)
}

func printStringList(res *list.List) {
	if nil == res {
		println("[]")
		return
	}

	print("[")
	for cur := res.Front(); nil != cur; {
		print(cur.Value.(string))
		cur = cur.Next()
		if nil != cur {
			print(", ")
		}
	}
	println("]")
}
