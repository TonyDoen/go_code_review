package chain

import "testing"

func TestDemoChain(t *testing.T) {
	//组装责任链
	h1 := NewConcreteHandler()
	h2 := NewConcreteHandler()
	h3 := NewConcreteHandler()

	h2.SetSuccessor(h3)
	h1.SetSuccessor(h2)

	//提交请求
	h1.HandleRequest()
}
