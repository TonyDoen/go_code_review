package easy

import (
	"fmt"
	"testing"
)

func TestMaximizeDistance2ClosestPerson2 (t *testing.T) {
	input := []int{1,0,0,0,1,0,1}
	res := MaximizeDistance2ClosestPerson2(input)
	fmt.Printf("%d\n", res)
}
