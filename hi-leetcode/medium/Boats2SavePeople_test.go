package medium

import (
	"fmt"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	people := []int{1, 2}
	limit := 3
	res := NumRescueBoats(people, limit)
	fmt.Printf("%d\n", res)
}
