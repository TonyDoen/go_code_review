package flyweight

import "testing"

func TestDemoFlyWeight(t *testing.T) {
	factory := NewFlyWeightFactory()

	f1 := factory.GetFlyWeight("f1")
	f1.Action("out2333")
}


