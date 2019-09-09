package medium

import (
	"fmt"
	"testing"
)

func TestStockSpannerNext(t *testing.T) {
	arr := []int{100, 80, 60, 70, 60, 75, 85}

	ss := new(StockSpanner)
	for _, v := range arr {
		t := ss.next(v)
		fmt.Printf("%d\n", t)
	}
}

func TestStockSpanner2Next(t *testing.T) {
	arr := []int{100, 80, 60, 70, 60, 75, 85}

	ss := new(StockSpanner2)
	for _, v := range arr {
		t := ss.next(v)
		fmt.Printf("%d\n", t)
	}
}
