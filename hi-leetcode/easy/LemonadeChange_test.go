package easy

import (
	"fmt"
	"testing"
)

func TestLemonadeChange2(t *testing.T) {
	bills := []int{5, 5, 5, 10, 20}
	tf := LemonadeChange2(bills)
	fmt.Printf("%t\n", tf)
}
