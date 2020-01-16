package proxy

import "testing"

func TestDemoProxy(t *testing.T) {
	// 静态代理
	proxy := NewSubjectProxy("iAmProxy")
	proxy.Action()

}
