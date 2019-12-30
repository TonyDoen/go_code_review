package easy

import (
	"fmt"
	"testing"
)

func TestCircularArrayLoop1(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := CircularArrayLoop1(arr)
	fmt.Printf("%v \n", res) // visit
}

func TestCircularArrayLoop2(t *testing.T) {
	arr := []int{2, -1, 1, 2, 2}
	res := CircularArrayLoop2(arr)
	fmt.Printf("%v \n", res) // visit
}
