package medium

import (
	"testing"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	arr := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	res := EvaluateReversePolishNotation(arr)
	println(res)
}
