package medium

import "testing"

func TestPushDominoes(t *testing.T) {
	src := ".L.R...LR..L.."
	res := PushDominoes(src)
	println(res)
}
