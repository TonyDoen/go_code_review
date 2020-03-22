package medium

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 3}
	cost := []int{1, 4, 2, 2}
	res := CanCompleteCircuit(gas, cost)
	fmt.Printf("%d\n", res)
}
