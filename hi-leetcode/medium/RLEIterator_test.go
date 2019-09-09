package medium

import (
	"fmt"
	"testing"
)

func TestNewRLEIteratorOneNext(t *testing.T) {
	arr := []int{3, 8, 0, 9, 2, 5}
	ro := NewRLEIteratorOne(arr)

	input := []int{2, 1, 1, 2}
	for _, v := range input {
		t := ro.next(v)
		fmt.Printf("%d\n", t)
	}
}
