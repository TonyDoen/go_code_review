package hard

import (
	"fmt"
	"testing"
)

func TestGiveCandy(t *testing.T) {
	arr := []int{4, 2, 3, 4, 1}
	res := GiveCandy(arr)
	fmt.Printf("%d\n", res)
}
