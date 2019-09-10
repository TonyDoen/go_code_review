package hard

import (
	"testing"
)

func TestFreqStack1(t *testing.T) {
	fs := NewFreqStack()

	fs.push(5)
	fs.push(7)
	fs.push(5)
	fs.push(7)
	fs.push(4)
	fs.push(5)
	fs.pop()
	fs.pop()
	fs.pop()
	fs.pop()
	fs.pop()
	fs.pop()
}
