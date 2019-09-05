package medium

import (
	"fmt"
	"testing"
)

func TestTotalFruit0(t *testing.T) {
	src := []int{0, 1, 2, 2}
	res := TotalFruit0(src)
	fmt.Printf("%d\n", res)
}

func TestTotalFruit2(t *testing.T) {
	src := []int{0, 1, 2, 2}
	res := TotalFruit2(src)
	fmt.Printf("%d\n", res)
}
