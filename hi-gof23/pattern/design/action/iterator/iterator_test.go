package iterator

import (
	"testing"
)

func TestDemoIterator(t *testing.T) {
	a := &ConcreteAggregate{container: []interface{}{"One", "Two", "Three", "Four", "Five", "Six"}}
	i := a.Iterator()
	for i.HasNext() {
		println(i.Next().(string))
	}
}
