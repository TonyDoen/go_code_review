package command

import "testing"

func TestDemoCommand(t *testing.T) {
	receiver := &Receiver{}
	cmd := &ConcreteCommand{r: receiver}
	invoker := &Invoker{c: cmd}
	invoker.action()
}
