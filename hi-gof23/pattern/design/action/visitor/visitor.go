package visitor

/**
 * 访问者模式把数据结构和作用于结构上的操作解耦合，使得操作集合可相对自由地演化。
 *
 * 访问者模式适用于数据结构相对稳定算法又易变化的系统。因为访问者模式使得算法操作增加变得容易。
 * 若系统数据结构对象易于变化，经常有新的数据对象增加进来，则不适合使用访问者模式。
 *
 * 访问者模式的优点是增加操作很容易，因为增加操作意味着增加新的访问者。访问者模式将有关行为集中到一个访问者对象中，其改变不影响系统数据结构。
 * 其缺点就是增加新的数据结构很困难。—— From 百科
 *
 * 简单来说，访问者模式就是一种分离对象数据结构与行为的方法，通过这种分离，可达到为一个被访问者动态添加新的操作而无需做其它的修改的效果。简单关系图：
 *
 *
 * 访问者模式的定义
 * 封装某些作用于某种数据结构中各元素的操作，它可以在不改变数据结构的前提下定义作用于这些元素的新的操作。
 *
 * 看一下在类A中，方法method1和方法method2的区别在哪里，方法method1很简单，就是打印出一句“我是A”；方法method2稍微复杂一点，使用类B作为参数，并调用类B的showA方法。再来看一下类B的showA方法，showA方法使用类A作为参数，然后调用类A的method1方法，可以看到，method2方法绕来绕去，无非就是调用了一下自己的method1方法而已，它的运行结果应该也是“我是A”，分析完之后，我们来运行一下这两个方法，
 *
 * 访问者模式的角色
 * 在例子中，对于类A来说，类B就是一个访问者。
 * 1. 抽象访问者：抽象类或者接口，声明访问者可以访问哪些元素，具体到程序中就是visit方法中的参数定义哪些对象是可以被访问的。
 * 2. 访问者：   实现抽象访问者所声明的方法，它影响到访问者访问到一个类后该干什么，要做什么事情。
 * 3. 抽象元素类：接口或者抽象类，声明接受哪一类访问者访问，程序上是通过accept方法中的参数来定义的。抽象元素一般有两类方法，一部分是本身的业务逻辑，另外就是允许接收哪类访问者来访问。
 * 4. 元素类：   实现抽象元素类所声明的accept方法，通常都是visitor.visit(this)，基本上已经形成一种定式了。
 * 5. 结构对象：  一个元素的容器，一般包含一个容纳多个不同类. 不同接口的容器，如List. Set. Map等，在项目中一般很少抽象出这个角色。
 *
 * 访问者模式的优点
 * 1. 符合单一职责原则：凡是适用访问者模式的场景中，元素类中需要封装在访问者中的操作必定是与元素类本身关系不大且是易变的操作，使用访问者模式一方面符合单一职责原则，另一方面，因为被封装的操作通常来说都是易变的，所以当发生变化时，就可以在不改变元素类本身的前提下，实现对变化部分的扩展。
 * 2. 扩展性良好：元素类可以通过接受不同的访问者来实现对不同操作的扩展。
 *
 * 访问者模式的缺点
 * 1. 增加新的元素类比较困难。通过访问者模式的代码可以看到，在访问者类中，每一个元素类都有它对应的处理方法，也就是说，每增加一个元素类都需要修改访问者类（也包括访问者类的子类或者实现类），修改起来相当麻烦。也就是说，在元素类数目不确定的情况下，应该慎用访问者模式。所以，访问者模式比较适用于对已有功能的重构，比如说，一个项目的基本功能已经确定下来，元素类的数据已经基本确定下来不会变了，会变的只是这些元素内的相关操作，这时候，我们可以使用访问者模式对原有的代码进行重构一遍，这样一来，就可以在不修改各个元素类的情况下，对原有功能进行修改。
 *
 * 访问者模式的适用场景
 * 1. 假如一个对象中存在着一些与本对象不相干（或者关系较弱）的操作，为了避免这些操作污染这个对象，则可以使用访问者模式来把这些操作封装到访问者中去。
 * 2. 假如一组对象中，存在着相似的操作，为了避免出现大量重复的代码，也可以将这些重复的操作封装到访问者中去。
 *
 * 该模式适用场景：
 * 如果我们想为一个现有的类增加新功能，不得不考虑几个事情：
 * 1. 新功能会不会与现有功能出现兼容性问题？
 * 2. 以后会不会再需要添加？
 * 3. 如果类不允许修改代码怎么办？
 *
 * 面对这些问题，最好的解决方法就是使用访问者模式，访问者模式适用于数据结构相对稳定的系统，把数据结构和算法解耦，
 *
 */

import "fmt"

// 访问者模式DEMO
type Visitor interface {
	Visit(Subject)
}

type Subject interface {
	Accept(Visitor)
	getSubject() Subject
}

type MyVisitor struct {
	name string
}

func (v MyVisitor) Visit(s Subject) {
	fmt.Print("v:")
	fmt.Print(v)
	fmt.Print(" has visit ")
	fmt.Print("s:")
	fmt.Println(s)
	fmt.Println()
}

type MySubject struct {
	name string
}

func (s MySubject) Accept(v Visitor) {
	fmt.Print("s:")
	fmt.Print(s)
	fmt.Print(" has accept ")
	fmt.Print("v:")
	fmt.Println(v)
	fmt.Println()
}

func (s MySubject) getSubject() Subject {
	return s
}

/**
 * 不确切的改造例子
 */
type EnvInfo1 struct {
}

func GetEnv() string {
	// 改造成读取环境变量
	return "test"
}

func (e EnvInfo1) Print() {
	if GetEnv() == "test" {
		fmt.Println("这是测试环境")
	}

	if GetEnv() == "production" {
		fmt.Println("这是生产环境")
	}
}

// 使用 访问者模式，改造一下
// 1. 定义访问者接口
type IVisitor interface {
	Visit() // 访问者的访问方法
}

// 2. 实现该接口
type ProductionVisitor struct {
}

func (v ProductionVisitor) Visit() {
	fmt.Println("这是生产环境")
}

type TestVisitor struct {
}

func (t TestVisitor) Visit() {
	fmt.Println("这是测试环境")
}

// 3. 创建元素接口
type IElement interface {
	Accept(visitor IVisitor)
}

// 4. 实现元素接口
type Element struct {
}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

// 5. 改造环境变量方法
type EnvInfo2 struct {
	Element
}

func (e EnvInfo2) Print(visitor IVisitor) {
	e.Element.Accept(visitor)
}
