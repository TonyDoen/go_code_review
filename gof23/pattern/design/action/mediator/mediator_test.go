package mediator

import (
	"testing"
)

func TestDemoMediator(t *testing.T) {
	mediator := MyMediator{}
	mediator.CreateMediator()
	mediator.WorkAll()
}

