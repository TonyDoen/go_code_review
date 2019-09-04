package easy

import (
	"fmt"
	"testing"
)

func TestUniqueMorseRepresentations1 (t *testing.T) {
	input := []string{"gin", "zen", "gig", "msg"}
	res := UniqueMorseRepresentations1(input)
	fmt.Printf("%d\n", res)
}

