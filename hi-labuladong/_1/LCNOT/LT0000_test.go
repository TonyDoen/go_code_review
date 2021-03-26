package LCNOT

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	var s1, s2 = "horse", "ros"
	var rs = MinDistance(s1, s2)
	fmt.Printf("%d\n", rs)
}
