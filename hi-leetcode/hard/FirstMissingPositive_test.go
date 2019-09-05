package hard

import (
	"fmt"
	"testing"
)

func TestGetFirstMissingPositive1(t *testing.T) {
	arr := []int{3, 4, -1, 1}
	res := GetFirstMissingPositive1(arr)
	fmt.Printf("%d\n", res)
}

func TestGetFirstMissingPositiveConstantSpace1(t *testing.T) {
	arr := []int{3, 4, -1, 1}
	res := GetFirstMissingPositiveConstantSpace1(arr)
	fmt.Printf("%d\n", res)
}
