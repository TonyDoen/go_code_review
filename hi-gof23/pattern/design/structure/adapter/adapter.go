package adapter

/**
 * 适配器模式类图
 *
 *                                _________________________________
 *  ___________________           |          <<ITarget>>           |
 * |       Client      |          |--------------------------------|
 * |-------------------|          |                                |
 * |                   |--------->|+request(): void                |
 * |-------------------|          |________________________________|
 * |                   |                                      /\
 * |___________________|                             _________|_________
 *                                                  |                   |
 *                            ______________________|_________     _____|__________________________
 *                           |       Adapter                  |   |       ConcreteTarget           |
 *                           |--------------------------------|   |--------------------------------|
 *  ___________________      |-adaptee: Adaptee               |   |                                |
 * |       Adaptee     |     |________________________________|   |________________________________|
 * |-------------------|     |+request(): void                |   |+request(): void                |
 * |                   |<----|________________________________|   |________________________________|
 * |-------------------|
 * |+onRequest(): void |
 * |___________________|
 *
 *
 * 适配器模式（Adapter Pattern），将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
 *
 * 适配器模式角色划分:
 * 1. 目标接口:     ITarget
 * 2. 具体目标实现:  ConcreteTarget
 * 3. 适配器:       Adapter
 * 4. 待适配类:     Adaptee
 *
 *
 * 适配器模式适用场景:
 * 1. 调用双方接口不一致且都不容易修改时，可以使用适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
 * 2. 多个组件功能类似，但接口不统一且可能会经常切换时，可使用适配器模式，使得客户端可以以统一的接口使用它们
 *
 *
 * 适配器模式优点:
 * 1. 客户端可以以统一的方式使用ConcreteTarget和Adaptee
 * 2. 适配器负责适配过程，而不需要修改待适配类，其它直接依赖于待适配类的调用方不受适配过程的影响
 * 3. 可以为不同的目标接口实现不同的适配器，而不需要修改待适配类，符合开放-关闭原则
 *
 * 适配器模式缺点
 *
 *
 * 已遵循的原则:
 * 1. 依赖倒置原则
 * 2. 迪米特法则
 * 3. 里氏替换原则
 * 4. 接口隔离原则
 * 5. 单一职责原则
 * 6. 开闭原则
 *
 *
 * 适配器模式可将一个类的接口转换成调用方希望的另一个接口。这种需求往往发生在后期维护阶段，因此有观点认为适配器模式只是前期系统接口设计缺乏的一种弥补。从实际工程来看，并不完全这样，有时不同产商的功能类似但接口很难完全一样，而为了系统使用方式的一致性，也会用到适配器模式。
 */

// 1. 目标接口
type ITarget interface {
	Request()
}

// 2. 具体目标实现
type ConcreteTarget struct {
	name string
}

func (ct ConcreteTarget) Request() {
	println(ct.name + " request now")
}

// 3. 适配器
type Adapter struct {
	ae Adaptee
}

func (a Adapter) Request() {
	println(" Adapter request now ")
	a.ae.FunkRequest()
}

func NewAdapter(name string, gender, role uint8) *Adapter {
	return &Adapter{ae: Adaptee{name: name, gender: gender, role: role}}
}

// 4. 待适配类
type Adaptee struct {
	name   string
	gender uint8
	role   uint8
}

func (a Adaptee) FunkRequest() {
	gender := "无性别记录"
	if 0 == a.gender {
		gender = "女"
	} else if 1 == a.gender {
		gender = "男"
	}
	println("被适配者 姓名:" + a.name + ", 性别:" + gender + " FunkRequest now")
}
