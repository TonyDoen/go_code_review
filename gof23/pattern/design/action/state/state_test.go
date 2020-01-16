package state

import "testing"

func TestDemoState(t *testing.T) {
	context := &Context{}
	context.SetState(&BlackState{})

	context.Pull()
	context.Push()
}
