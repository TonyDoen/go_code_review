package easy

import (
	"fmt"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	arr := []int{1, 2, 2, 3}
	res := IsMonotonic(arr)
	fmt.Printf("%v \n", res) // visit
}

func TestIsMonotonic2(t *testing.T) {
	arr := []int{1, 2, 2, 3}
	res := IsMonotonic2(arr)
	fmt.Printf("%v \n", res) // visit
}

