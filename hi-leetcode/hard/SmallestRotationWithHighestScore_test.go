package hard

import (
	"fmt"
	"testing"
)

func TestBestRotation1(t *testing.T)  {
	arr := []int{2, 3, 1, 4, 0}
	res := BestRotation1(arr)
	fmt.Printf("%d\n", res)
}
