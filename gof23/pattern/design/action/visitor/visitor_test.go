package visitor

import (
	"testing"
)

func TestEnvAlter(t *testing.T) {
	e := new(Element)
	e.Accept(new(ProductionVisitor)) // output: 这是生产环境
	e.Accept(new(TestVisitor))       // output: 这是测试环境

	m := new(EnvInfo2)
	m.Print(new(ProductionVisitor))

}

func TestDemoVisitor(t *testing.T) {
	//visitor := new(MyVisitor)
	//subject := new(MySubject)
	visitor := &MyVisitor{name: "myVisitor11"}
	subject := &MySubject{name: "mySubject11"}
	subject.Accept(visitor)
}