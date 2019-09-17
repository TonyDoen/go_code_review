package hard

import (
	"fmt"
	"testing"
)

func TestAtMostNGivenDigitSet1(t *testing.T) {
	set := NewSetWithItem("1","3","5","7")
	res := AtMostNGivenDigitSet(set, 365)
	fmt.Println(res)
}
