package medium

import (
	"fmt"
	"testing"
)

func TestSurroundedRegions(t *testing.T) {
	matrix := [][]rune{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'}}
	SurroundedRegions(matrix)

	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		println()
	}
	println()
}
