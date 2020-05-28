package interpreter

/**
 * 解释器模式（Interpreter）
 *
 * 解释器模式是我们暂时的最后一讲，一般主要应用在OOP开发中的编译器的开发中，所以适用面比较窄。
 *
 * Context类是一个上下文环境类，Plus和Minus分别是用来计算的实现，
 *
 *
 */

// 1. 抽象表达式
type Expression interface {
	interpret(c Context) float64
}

// 2. 上下文环境类
type Context struct {
	_1, _2 float64
}

//func (c *Context)get1() float64 {
//	return c._1
//}
//
//func (c *Context)get2() float64 {
//	return c._2
//}

func NewContext(n1, n2 float64) *Context {
	return &Context{n1, n2}
}

// 3. 具体表达式
type Minus struct {
}

func (m *Minus) interpret(c *Context) float64 {
	return c._1 - c._2
}

type Plus struct {
}

func (p *Plus) interpret(c *Context) float64 {
	return c._1 + c._2
}
