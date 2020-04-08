package medium

import "testing"

func TestNewTicTacToe(t *testing.T) {
	n := 5
	ticTacToe := NewTicTacToe(n)
	for i := 0; i < n; i++ {
		blackResult := ticTacToe.Move(0, i, BlackPlayer)
		if 1 == blackResult {
			println("black win")
			return
		}
		redResult := ticTacToe.Move(3, i, RedPlayer)
		if 1 == redResult {
			println("red win")
			return
		}
	}
}
