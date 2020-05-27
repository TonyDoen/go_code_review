package hard

import (
	"fmt"
	"testing"
)

func TestMove2Stamp(t *testing.T) {
	stamp, target := "abc", "ababc"
	res := move2Stamp(stamp, target)
	for _, i := range res {
		fmt.Printf("%d ", i)
	}
}
