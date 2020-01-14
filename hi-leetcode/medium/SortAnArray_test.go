package medium

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	src := []int{5, 1, 1, 2, 0, 0}
	res := CountSort(src)
	for _, v := range res {
		fmt.Print(v, " ")
	}
}

func TestQuickSort(t *testing.T) {
	src := []int{5, 1, 1, 2, 0, 0}
	res := QuickSort(src)
	for _, v := range res {
		fmt.Print(v, " ")
	}
}
