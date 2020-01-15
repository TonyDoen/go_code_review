package template

import "strconv"

/**
 * 其实golang的结构体组合就是一种模板方法模式的实现，
 * golang不像JAVA等其他语言，缺失了一些特性，无法做到强制子类实现某些方法，
 * 但是上面的那种形式，确实符合模板方法模式的定义，规定了逻辑流程，那么我们就可以说是实现了模板方法模式，
 * 设计模式更多的是思想，在不同语言中都有不同的体现。
 *
 * 模板方法模式定义了一个演算法的步骤，并允许子类别为一个或多个步骤提供其实践方式。
 * 让子类别在不改变演算法架构的情况下，重新定义演算法中的某些步骤。
 * 在软件工程中，它是一种软件设计模式，和C++模板没有关连。
 * 简单一点说模板方法模式就是将一个类种能够公共使用的方法放置在抽象类种实现，不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要去填充的地方。
 *
 *
 *
 * golang并没有严格意义上的继承机制, 它仅仅是利用组合的特性来模拟继承, 在这方法对于组合优于继承体现的还是很好的
 *
 *
 *
 * 模板方法模式类图
 *                                           ____________________
 *                                          |    SchoolPerson    |
 *                                          |--------------------|
 *                                          |                    |
 *                                          |--------------------|
 *                                          |+prepareGotoSchool()|
 *                                          |-dressUp()          |
 *                                          |-eatBreakfast()     |
 *                                          |-takeThings()       |
 *                                          |____________________|
 *                                                    /\
 *                                                    |
 *                                           _  _  _ _|_ _ _ _ _
 *                         _________________|_____         _____|_________________
 *                        |         Teacher       |       |          Student      |
 *                        |-----------------------|       |-----------------------|
 *                        |_______________________|       |_______________________|
 *                        |                       |       |                       |
 *                        |-dressUp()             |       |-dressUp()             |
 *                        |-eatBreakfast()        |       |-eatBreakfast()        |
 *                        |-takeThings()          |       |-takeThings()          |
 *                        |_______________________|       |_______________________|
 *
 * 模板方法模式是类的行为模式。
 * 准备一个抽象类，将部分逻辑以具体方法以及具体构造函数的形式实现，
 * 然后声明一些抽象方法来迫使子类实现剩余的逻辑。
 * 不同的子类可以以不同的方式实现这些抽象方法，从而对剩余的逻辑有不同的实现。
 * 这就是模板方法模式的用意。
 *
 * 比如定义一个操作中的算法的骨架，将步骤延迟到子类中。模板方法使得子类能够不去改变一个算法的结构即可重定义算法的某些特定步骤。
 *
 * 模板方法模式，就是指：一个抽象类中，有一个主方法，再定义1...n个方法，可以是抽象的，也可以是实际的方法，定义一个类，继承该抽象类，重写抽象方法，通过调用抽象类，实现对子类的调用
 *
 *
 *
 *
 * 模板方法模式中的角色:
 * 1. 抽象类（AbstractClass）：实现了模板方法，定义了算法的骨架。
 * 2. 具体类（ConcreteClass)：实现抽象类中的抽象方法，已完成完整的算法
 *
 * 模板方法模式的优点:
 * 1. 模板方法模式通过把不变的行为搬移到超类，去除了子类中的重复代码。
 * 2. 子类实现算法的某些细节，有助于算法的扩展。
 * 3. 通过一个父类调用子类实现的操作，通过子类扩展增加新的行为，符合“开放-封闭原则”。
 * 3.1. 封装不变部分，扩展可变部分。
 * 3.2. 提取公共代码，便于维护。
 * 3.3. 行为由父类控制，子类实现。
 *
 * 模板方法模式的缺点
 * 1. 每个不同的实现都需要定义一个子类，这会导致类的个数的增加，设计更加抽象。(每一个不同的实现都需要一个子类来实现，导致类的个数增加，使得系统更加庞大)
 *
 * 模板方法模式的适用场景:
 * 1. 在某些类的算法中，用了相同的方法，造成代码的重复。
 * 2. 控制子类扩展，子类必须遵守算法规则。
 *
 * 2.1. 有多个子类共有的方法，且逻辑相同
 * 2.2. 重要的、复杂的方法，可以考虑作为模板方法
 *
 */

type ISchoolPerson interface {
	DressUp()
	EatBreakfast()
	TakeThings()
}

type SchoolPerson struct {
	one      ISchoolPerson
	name     string
	gender   int8
	birthday string
}

func (sp SchoolPerson) DressUp() {
	sp.one.DressUp()
}

func (sp SchoolPerson) EatBreakfast() {
	sp.one.EatBreakfast()
}

func (sp SchoolPerson) TakeThings() {
	sp.one.TakeThings()
}

// 算法骨架
func (sp SchoolPerson) PrepareGotoSchool() {
	if nil == sp.one {
		return
	}
	sp.DressUp()
	sp.EatBreakfast()
	sp.TakeThings()
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 去学校")
}

// Student
type Student struct {
	SchoolPerson // 组合优于继承; 提倡的是把所有逻辑都显式操作，所以没有继承、抽象等等东西，虽然可以使用结构体组合达到一些重写的效果，但是还是有一点区别的:
}

func (sp Student) DressUp() {
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 穿校服")
}

func (sp Student) EatBreakfast() {
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 吃妈妈做好的早饭")
}

func (sp Student) TakeThings() {
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 背书包，带上家庭作业和红领巾")
}

func NewStudent(name string, gender int8, birthday string) *Student {
	s := new(Student)
	s.SchoolPerson = SchoolPerson{one: s, name: name, gender: gender, birthday: birthday}
	return s
}

// Teacher
type Teacher struct {
	SchoolPerson
}

func (sp Teacher) DressUp() {
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 穿工作服")
}

func (sp Teacher) EatBreakfast() {
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 做早饭，照顾孩子吃早饭")
}

func (sp Teacher) TakeThings() {
	println("{name:" + sp.name + ", gender:" + strconv.FormatInt(int64(sp.gender), 10) + ", birthday:" + sp.birthday + "} 带上昨晚准备的考卷")
}

func NewTeacher(name string, gender int8, birthday string) *Teacher {
	t := new(Teacher)
	t.SchoolPerson = SchoolPerson{one: t, name: name, gender: gender, birthday: birthday}
	return t
}
