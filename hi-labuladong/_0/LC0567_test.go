package _0

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(tg *testing.T) {
	var t, s = "ab", "eidbaooo"
	var rs = CheckInclusion(t, s)
	fmt.Printf("%v\n", rs)
}
