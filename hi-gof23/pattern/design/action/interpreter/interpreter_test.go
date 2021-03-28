package interpreter

import "testing"

func TestDemoInterpreter(t *testing.T) {
	plus := &Plus{}
	minus := &Minus{}

	c1 := NewContext(9, 2)
	res1 := plus.interpret(c1)

	c2 := NewContext(res1, 8)
	res2 := minus.interpret(c2)

	println(res1)
	println(res2)
}
