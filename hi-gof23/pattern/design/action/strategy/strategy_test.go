package strategy

import (
	"testing"
)

func TestDemoStrategy(t *testing.T) {
	ctx := Context{strategy: BubbleSortStrategy{}}
	ctx.Sort()

	ctx = Context{strategy: MergeSortStrategy{}}
	ctx.Sort()
}
