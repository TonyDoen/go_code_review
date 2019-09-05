package easy

import (
	"fmt"
	"testing"
)

func TestBackspaceCompare1(t *testing.T) {
	tf := BackspaceCompare1("ab#c", "ad#c")
	fmt.Printf("%t\n", tf)
}

func TestBackspaceCompare2(t *testing.T) {
	tf := BackspaceCompare2("ab#c", "ad#c")
	fmt.Printf("%t\n", tf)
}

func TestBackspaceCompare2Nil(t *testing.T) {
	var p string
	tf := BackspaceCompare2(p, p)
	fmt.Printf("%t\n", tf)
}
