package proxy

/**
 * 代理模式类图                     _________________________________
 *                                |          ISubject              |
 *                                |--------------------------------|
 *  ___________________           |+action(): void                 |
 * |       Client      |          |                                |
 * |-------------------|          |                                |
 * |                   |--------->|                                |
 * |-------------------|          |________________________________|
 * |                   |                       /\
 * |___________________|              _ _ _ _ _|_ _ _ _ _
 *                                   |                   |
 *             ______________________|_________       ___|____________________________
 *            |       ConcreteSubject          |     |       SubjectProxy             |
 *            |--------------------------------|     |--------------------------------|
 *            |                                |     |-target: ISubject               |
 *            |________________________________|     |________________________________|
 *            |+action() :void                 |     |-preAction(): void              |
 *            |                                |<--<>|-postAction(): void             |
 *            |                                |     |+action(): void                 |
 *            |                                |     |                                |
 *            |________________________________|     |________________________________|
 *
 * 代理模式（Proxy Pattern），为其它对象提供一种代理以控制对这个对象的访问。
 * 静态代理
 * 动态代理
 *
 * 从上述代码中可以看到，被代理对象由代理对象在编译时确定，并且代理对象可能限制对被代理对象的访问。
 * 调用方直接调用代理而不需要直接操作被代理对象甚至都不需要知道被代理对象的存在。同时，代理类可代理的具体被代理类是确定的，如本例中ProxySubject只可代理ConcreteSubject。
 */

// 1. 被代理抽象类
type ISubject interface {
	Action()
}

// 2. 被代理具体类(被装饰抽象类实现)
type ConcreteSubject struct {
	name string
}

func (cs *ConcreteSubject) Action() {
	println("ConcreteSubject Name:" + cs.name + " Action()")
}

// 3. 具体代理类(代理类实现)
type SubjectProxy struct {
	sub ISubject
}

func (s *SubjectProxy) Action() {
	s.beforeAction()
	s.sub.Action()
	s.postAction()
}

func (s *SubjectProxy) postAction() {
	println("after ISubject. SubjectProxy postAction()")
}

func (s *SubjectProxy) beforeAction() {
	println("before ISubject. SubjectProxy beforeAction()")
}

func NewSubjectProxy(name string) *SubjectProxy {
	// 静态代理
	return &SubjectProxy{&ConcreteSubject{name: name}} // 此处就是 装饰模式 与 代理模式的区别
}
