package decorator

import "testing"

func TestDemoDecorator(t *testing.T) {
	target := NewConcreteSubject("666Target")  // 此处就是 装饰模式 与 代理模式的区别
	postDecorator := NewSubjectPostDecorator(target)
	postDecorator.Action()
}

