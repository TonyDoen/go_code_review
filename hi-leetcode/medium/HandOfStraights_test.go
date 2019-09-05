package medium

import (
	"fmt"
	"testing"
)

func TestIsStraightHand1(t *testing.T) {
	src := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	w := 3
	res := IsStraightHand1(src, w)
	fmt.Printf("%v\n", res)
}
