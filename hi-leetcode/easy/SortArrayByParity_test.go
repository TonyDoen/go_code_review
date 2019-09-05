package easy

import (
	"fmt"
	"testing"
)

func TestSortArrayByParity1(t *testing.T) {
	input := []int{3, 1, 2, 4}
	//res := SortArrayByParity1(input)
	res := SortArrayByParity2(input)
	for _, i := range res {
		fmt.Printf("%d ", i)
	}
}
