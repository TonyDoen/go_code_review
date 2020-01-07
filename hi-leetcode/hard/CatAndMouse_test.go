package hard

import (
	"fmt"
	"testing"
)

func TestCatAndMouseGame1(t *testing.T) {
	arr := [][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}
	res := CatAndMouseGame1(arr)
	fmt.Printf("%d\n", res)
}
