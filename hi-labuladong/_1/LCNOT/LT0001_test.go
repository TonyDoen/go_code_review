package LCNOT

import (
	"fmt"
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	var envelopes = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	var rs = MaxEnvelopes(envelopes)
	fmt.Printf("%d\n", rs)
}
