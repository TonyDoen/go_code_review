package _0

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(tg *testing.T) {
	var s, t = "cbaebabacdabc", "abc"
	var rs = FindAnagrams(s, t)

	for tmp := rs.Front(); nil != tmp; tmp = tmp.Next() {
		fmt.Printf("%d, ", tmp.Value.(int))
	}
}
