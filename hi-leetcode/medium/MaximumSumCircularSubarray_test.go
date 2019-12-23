package medium

import (
	"fmt"
	"testing"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	arr := []int{100, 80, 60, 70, 60, 75, 85}

	res := maxSubarraySumCircular(arr)
	fmt.Printf("%d\n", res)
}
