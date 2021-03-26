package _0

import (
	"fmt"
	"testing"
)

func TestMinWindow(tg *testing.T) {
	var s, t = "BANCADBCEFGDAFSDFE", "ABC"
	var rs = MinWindow(s, t)
	fmt.Printf("%s\n", rs)
}
