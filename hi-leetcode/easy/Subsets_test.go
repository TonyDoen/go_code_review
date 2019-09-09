package easy

import (
	"fmt"
	"testing"
)

func TestSubset1(t *testing.T) {
	input := []int{1, 2, 3}
	res := Subset1(input)
	fmt.Println(res)
}

func TestSubset2(t *testing.T) {
	input := []int{1, 2, 3}
	res := Subset2(input)
	fmt.Println(res)
}
