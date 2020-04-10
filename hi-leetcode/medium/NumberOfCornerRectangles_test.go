package medium

import "testing"

func TestCountCornerRectangles(t *testing.T) {
	grid := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	res := CountCornerRectangles(grid)
	println(res)
}

func TestCountCornerRectangles1(t *testing.T) {
	grid := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	res := CountCornerRectangles1(grid)
	println(res)
}
