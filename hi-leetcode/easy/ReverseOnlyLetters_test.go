package easy

import (
	"fmt"
	"testing"
)

func TestReverseOnlyLetters1(t *testing.T) {
	src := "a-bC-dEf-ghIj"

	res := ReverseOnlyLetters1(src)
	fmt.Printf("%s \n", res) // visit
}

func TestReverseOnlyLetters2(t *testing.T) {
	src := "a-bC-dEf-ghIj"

	res := ReverseOnlyLetters2(src)
	fmt.Printf("%s \n", res) // visit
}

