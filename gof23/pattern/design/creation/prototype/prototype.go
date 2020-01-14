package prototype

/**
 * 原型模式: 用原型实例指定创建对象的种类，并通过拷贝这些原型创建新的对象。
 *
 * 原型模式主要用于对象的复制，它的核心是就是类图中的原型类Prototype。
 * Prototype类需要具备以下两个条件：
 * java-1. 实现Cloneable接口。在java语言有一个Cloneable接口，它的作用只有一个，就是在运行时通知虚拟机可以安全地在实现了此接口的类上使用clone方法。在java虚拟机中，只有实现了这个接口的类才可以被拷贝，否则在运行时会抛出CloneNotSupportedException异常。
 * java-2. 重写Object类中的clone方法。Java中，所有类的父类都是Object类，Object类中有一个clone方法，作用是返回对象的一个拷贝，但是其作用域protected类型的，一般的类无法调用，因此，Prototype类需要将clone方法的作用域修改为public类型。
 *
 * 原型模式是一种比较简单的模式，也非常容易理解，实现一个接口，重写一个方法即完成了原型模式。在实际应用中，原型模式很少单独出现。经常与其他模式混用，他的原型类Prototype也常用抽象类来替代。
 *
 *
 * 原型模式的优点：
 * java-1. 使用原型模式创建对象比直接 new 一个对象在性能上要好的多，因为Object类的clone方法是一个本地方法，它直接操作内存中的二进制流，特别是复制大对象时，性能的差别非常明显。
 * java-2. 使用原型模式的另一个好处是简化对象的创建，使得创建对象就像我们在编辑文档时的复制粘贴一样简单。
 *
 * 1. 对客户隐藏了具体的产品类，因此减少了客户知道对象的数目。此外，这些模式使客户无需改变即可使用与特定应用相关的类（和工厂模式. 建造者模式一样的效果）
 * 2. 运行时刻增加和删除产品，原型模式允许只通过客户注册原型实例就可以将一个新的具体产品类型注入系统。相比其他创建型模式更为灵活，因此客户可以在运行时刻建立和删除原型。（可能开发时候方便，要注意可维护性）
 * 3. 改变值以指定新对象，高度动态的系统允许通过对象复合定义新的行为。例如，通过一个对象变量指定值—-并不定义新的类型。通过实例化已有类并且将这些实例注册为客户对象的原型，就可以有效定义新类别的对象。客户可以将责职代理给原型，从而表现出新的行为。
 * 4. 改变结构以指定新对象。
 * 5. 减少子类的构造。这个在适用性中已经介绍过。
 * 6. 用类型动态配置应用，一些运行时刻环境允许动态将类装载到应用中。
 *
 * 原型模式的缺点：
 * 1. 每种具体实现类型都要有一个克隆自己的操作。在某些场景会比较困难。
 *
 *
 * 原型模式的使用场景：
 * 1. 当一个系统应该独立于它的产品创建. 构成和表示时，要使用原型模式；
 * 2. 当要实例化的类是在运行时刻指定时，例如动态装载；
 * 3. 为了避免创建一个与产品类层次平行的工厂类层次时（和下一点很像，系统初始创建一个相对复杂的对象，也可以用在这基础上实现工厂模式，但是创建出来的对象和原来的对象在同一个层次）；
 * 4. 当一个类的实例只能有几个不同状态组合中的一种时，建立相应数目的原型并克隆它们可能比每次用合适的状态手工实例化对象更方便些。各种实例间的差异很小。但是初始化过程可能比较复杂。
 *
 *
 * 原型模式的注意事项：
 * java-1. 使用原型模式复制对象不会调用类的构造方法。因为对象的复制是通过调用 Object类的clone方法来完成的，它直接在内存中复制数据，因此不会调用到类的构造方法。不但构造方法中的代码不会执行，甚至连访问权限都对原型模式无效。单例模式中，只要将构造方法的访问权限设置为private型，就可以实现单例。但是clone方法直接无视构造方法的权限，所以，单例模式与原型模式是冲突的，在使用时要特别注意。
 * java-2. 深拷贝与浅拷贝。Object类的clone方法只会拷贝对象中的基本的数据类型（8种基本数据类型byte,char,short,int,long,float,double,boolean），对于数组. 容器对象. 引用对象等都不会拷贝，这就是浅拷贝。如果要实现深拷贝，必须将原型模式中的数组. 容器对象. 引用对象等另行拷贝。
 * java-3. 关于深拷贝和浅拷贝，会发生深拷贝的是java 的 8种基本数据类型和他们的封装类，至于String这个类型需要注意，它是引用数据类型，所以是浅拷贝
 *
 *
 * 从例子中可以看到如果重新去创建ThisWriter这个对象，就会像新建role一样，这里类型相对简单，在实际开发中可能会是个比较复杂的类型。
 * 但是role和ThisWriter其他都一致，只有Tall不一致，如果重新创建会比较繁琐。
 * 用工厂类的话也会不是很合适，因为role和ThisWriter只有Tall属性不一致。
 * 如果再增加几个不同小改动的Role，使用工厂类会引入大量的类。所以使用原型模式最为合适。
 */

import "fmt"

type Role interface {
	Clone() Role
	SetHeadColor(string)
	SetEyesColor(string)
	SetTall(int64)
	Show()
}

type RoleChinese struct {
	HeadColor string
	EyesColor string
	Tall      int64
}

func (pR *RoleChinese) Clone() Role {
	var pChinese = &RoleChinese{HeadColor: pR.HeadColor, EyesColor: pR.EyesColor, Tall: pR.Tall}
	return pChinese
}

func (pR *RoleChinese) SetHeadColor(color string) {
	pR.HeadColor = color
}

func (pR *RoleChinese) SetEyesColor(color string) {
	pR.EyesColor = color
}

func (pR *RoleChinese) SetTall(tall int64) {
	pR.Tall = tall
}

func (pR *RoleChinese) Show() {
	fmt.Println("Role's headColor is:", pR.HeadColor, " EyesColor is:", pR.EyesColor, " tall is:", pR.Tall)
}
