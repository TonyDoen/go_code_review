package medium

import (
	"fmt"
	"testing"
)

func TestSubarrayBitwiseORs(t *testing.T) {
	arr := []int{0, 3, 4, 6, 5}
	res := SubarrayBitwiseORs(arr)
	fmt.Printf("%d\n", res)
}
