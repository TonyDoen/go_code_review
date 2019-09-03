package easy

import (
	"fmt"
	"testing"
)

func TestLetterCasePermutation4 (t *testing.T) {
	input := "a1b2"
	res := LetterCasePermutation4(input)
	for e := res.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s\n", e.Value)
	}
}
