package adapter

import (
	"testing"
)

func TestDemoAdapter(t *testing.T) {
	ad := NewAdapter("tony", 0, 0)
	ad.Request()

	ct := ConcreteTarget{"tommy"}
	ct.Request()
}
