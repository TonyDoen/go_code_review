package easy

import (
	"fmt"
	"testing"
)

func TestHasGroupsSizeX1(t *testing.T) {
	deck := []int{1, 2, 3, 4, 4, 3, 2, 1}
	res := HasGroupsSizeX1(deck)
	fmt.Println(res)
}

func TestHasGroupsSizeX2(t *testing.T) {
	deck := []int{1, 2, 3, 4, 4, 3, 2, 1}
	res := HasGroupsSizeX2(deck)
	fmt.Println(res)
}
