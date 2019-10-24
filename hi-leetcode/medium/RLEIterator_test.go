package medium

import (
	"fmt"
	"testing"
)

func TestNewRLEIteratorOneNext(t *testing.T) {
	arr := []int{3, 8, 0, 9, 2, 5}
	rr := NewRLEIterator1(arr)

	input := []int{2, 1, 1, 2}
	for _, v := range input {
		t := rr.next(v)
		fmt.Printf("%d\n", t)
	}

	rr2 := NewRLEIterator2(arr)
	for _, v := range input {
		t := rr2.next(v)
		fmt.Printf("%d\n", t)
	}
}
