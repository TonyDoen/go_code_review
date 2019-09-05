package medium

import (
	"fmt"
	"testing"
)

func TestShiftingLetters1(t *testing.T) {
	src := "abc"
	shifts := []int{3, 5, 9}
	res := ShiftingLetters1(src, shifts)
	fmt.Printf("%s\n", res)
}
