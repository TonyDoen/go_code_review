package medium

import (
	"fmt"
	"testing"
)

func TestWordSubset1(t *testing.T) {
	// A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
	a, b := []string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}
	res := WordSubset1(a, b)

	for e := res.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		fmt.Print(", ")
	}
}
