package medium

import (
	"fmt"
	"testing"
)

func TestPartitionDisjoint1(t *testing.T) {
	arr := []int{5, 0, 3, 8, 6}
	res := PartitionDisjoint1(arr)
	fmt.Println(res)
}

func TestPartitionDisjoint2(t *testing.T) {
	arr := []int{5, 0, 3, 8, 6}
	res := PartitionDisjoint2(arr)
	fmt.Println(res)
}
