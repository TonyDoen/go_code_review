package observer

import "testing"

func TestDemoObserver(t *testing.T) {
	s := new(Subject)  // 创建被观察者

	o := new(Observer) // 创建观察者
	s.AddObserver(o)   // 为主题添加观察者

	o1 := new(Observer)
	s.AddObserver(o1)

	// 被观察者要做各种更改 ...

	// 更改完毕，触发观察者
	s.NotifyObserver()
}
