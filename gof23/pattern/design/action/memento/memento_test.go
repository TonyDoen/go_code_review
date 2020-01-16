package memento

import (
	"fmt"
	"testing"
)

func TestDemoStudent(t *testing.T) {

	// 创建一个发起人并设置初始状态, 此时与备忘录模式无关，只是模拟正常程序运行
	o := &Original{state: "hello"}

	// 现在需要保存当前状态, 就创建一个负责人来设置（一般来说，对于一个对象的同一个备忘范围，应当只有一个负责人，这样方便做多状态多备忘管理）
	c := new(Storage)
	c.SetMemento(o.CreateMemento())

	// 当前状态
	fmt.Println("当前状态:", o.GetState())

	// 改变状态
	o.SetState("world")
	fmt.Println("更改当前状态:", o.GetState())

	// 恢复备忘
	o.SetState(c.GetMemento().GetState())
	fmt.Println("恢复后状态", o.GetState())
}
