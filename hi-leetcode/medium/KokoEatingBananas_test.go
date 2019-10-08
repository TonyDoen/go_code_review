package medium

import (
	"fmt"
	"testing"
)

func TestMinEatingSpeed1(t *testing.T) {
	var piles, h = []int{3, 6, 7, 11}, 4
	res := MinEatingSpeed1(piles, h)
	fmt.Printf("%v\n", res)
}
