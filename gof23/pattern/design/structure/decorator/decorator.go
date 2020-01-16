package decorator

/**
 * 装饰模式类图                     _________________________________
 *                                |          ISubject              |
 *                                |--------------------------------|
 *  ___________________           |+action(): void                 |
 * |       Client      |          |                                |
 * |-------------------|          |                                |<------------------------
 * |                   |--------->|                                |                         |
 * |-------------------|          |________________________________|                         |
 * |                   |                       /\                                            |
 * |___________________|              _ _ _ _ _|_ _ _ _ _                                    |
 *                                   |                   |                                   |
 *             ______________________|_________       ___|____________________________       |
 *            |       ConcreteSubject          |     |       SubjectPostDecorator     |      |
 *            |--------------------------------|     |--------------------------------|      |
 *            |                                |     |-target: ISubject               |      |
 *            |________________________________|     |________________________________|      |
 *            |+action() :void                 |     |                                |      |
 *            |                                |     |-postAction(): void             |<>-----
 *            |                                |     |+action(): void                 |
 *            |                                |     |                                |
 *            |________________________________|     |________________________________|
 *
 * 装饰类可装饰的类并不固定，并且被装饰对象是在使用时通过组合确定。
 * 如本例中SubjectPreDecorator装饰ConcreteSubject，而SubjectPostDecorator装饰SubjectPreDecorator。并且被装饰对象由调用方实例化后通过构造方法（或者setter）指定。
 *
 * 装饰器本质上允许您包装现有功能并在开始或结尾处添加您自己的自定义功能。
 *
 * 装饰模式的本质是动态组合。动态是手段，组合是目的。
 * 每个装饰类可以只负责添加一项额外功能，然后通过组合为被装饰类添加复杂功能。
 * 由于每个装饰类的职责比较简单单一，增加了这些装饰类的可重用性，同时也更符合单一职责原则
 *
 *
 * 从语意上讲，
 * 代理模式是为控制对被代理对象的访问，
 * 而装饰模式是为了增加被装饰对象的功能
 *
 * 代理类所能代理的类完全由代理类确定，
 * 装饰类装饰的对象需要根据实际使用时客户端的组合来确定
 *
 * 被代理对象由代理对象创建，客户端甚至不需要知道被代理类的存在；
 * 被装饰对象由客户端创建并传给装饰对象
 */

// 1. 被装饰抽象类
type ISubject interface {
	Action()
}

// 2. 被装饰具体类(被装饰抽象类实现)
type ConcreteSubject struct {
	name string
}

func (cs *ConcreteSubject)Action()  {
	println("ConcreteSubject Name:" + cs.name + " action()")
}

func NewConcreteSubject(name string) *ConcreteSubject {
	return &ConcreteSubject{name: name}
}

// 3. 装饰类(装饰类实现)
type SubjectPostDecorator struct {
	sub ISubject
}

func (s *SubjectPostDecorator)Action()  {
	s.sub.Action()
	s.postAction()
}

func (s *SubjectPostDecorator)postAction()  {
	println("SubjectPostDecorator postAction()")
}

func NewSubjectPostDecorator(sub ISubject) *SubjectPostDecorator {
	return &SubjectPostDecorator{sub: sub} // 此处就是 装饰模式 与 代理模式的区别
}

