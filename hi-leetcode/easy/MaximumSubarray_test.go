package easy

import (
	"fmt"
	"testing"
)

func TestMaxSubArray1(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := MaxSubArray1(arr)
	fmt.Printf("%d \n", res) // visit
}

func TestMaxSubArray2(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := MaxSubArray2(arr)
	fmt.Printf("%d \n", res) // visit
}