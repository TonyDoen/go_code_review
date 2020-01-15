package observer

import "fmt"

/**
 * 观察者模式定义
 * 观察者模式又叫发布-订阅模式，它定义了一种一对多的依赖关系，多个观察者对象可同时监听某一主题对象，当该主题对象状态发生变化时，相应的所有观察者对象都可收到通知。
 *
 * 观察者模式角色划分
 * 1. 抽象被观察者(主题，抽象类或接口):  如上面类图中的 ISubject
 * 2. 具体被观察者(具体主题):          如上面类图中的 Subject, ...
 * 3. 抽象观察者(观察者):             如上面类图中的 IObserver
 * 4. 具体观察者:                    如上面类图中的 Observer, ...
 *
 *
 * 观察者模式优点
 * 1. 抽象主题只依赖于抽象观察者
 * 2. 观察者模式支持广播通信
 * 3. 观察者模式使信息产生层和响应层分离
 *
 * 观察者模式缺点
 * 1. 如一个主题被大量观察者注册，则通知所有观察者会花费较高代价
 * 2. 如果某些观察者的响应方法被阻塞，整个通知过程即被阻塞，其它观察者不能及时被通知
 *
 *
 * 已遵循的原则
 * 1. 依赖倒置原则（主题类依赖于抽象观察者而非具体观察者）
 * 2. 迪米特法则
 * 3. 里氏替换原则
 * 4. 接口隔离原则
 * 5. 单一职责原则
 * 6. 开闭原则
 *
 *
 * 观察者模式是软体设计模式的一种。
 * 在此种模式中，一个目标物件管理所有相依于它的观察者物件，并且在它本身的状态改变时主动发出通知。
 * 这通常透过呼叫各观察者所提供的方法来实现。此种模式通常被用来实时事件处理系统。
 * 简单点说，可以想象成多个对象同时观察一个对象，当这个被观察的对象发生变化的时候，这些对象都会得到通知，可以做一些操作...
 */

// 1. 创建抽象观察者
type IObserver interface {
	Notify() // 当被观察对象有更改的时候，出发观察者的Notify() 方法
}

// 2. 创建抽象被观察者
type ISubject interface {
	AddObserver(observers ...IObserver) // 添加观察者
	NotifyObserver()                    // 通知观察者
}

// 3. 实现具体观察者
type Observer struct {
}

func (o *Observer) Notify() {
	fmt.Println("已经触发了观察者")
}

// 4. 实现具体被观察者
type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObserver(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject) NotifyObserver() {
	for _, v := range s.observers {
		v.Notify() // 触发观察者
	}
}
