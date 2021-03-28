package composite

import "reflect"

/**
 * 组合模式类图                      _________________________________________
 *                                |          Organization                  |
 *                                |----------------------------------------|
 *                                |+GetChildOrganization(): []Organization |
 *                                |+GetName(): string                      |
 *  ___________________           |+Inform(info: String)                   |
 * |       Client      |          |+Remove(o: Organization)                |
 * |-------------------|          |+Add(o: Organization)                   |
 * |                   |--------->|                                        |<---------------------------|
 * |-------------------|          |________________________________________|                            |
 * |                   |                       /\                                                       |
 * |___________________|              _________|________________                                        |
 *                                   |                          |                                       |
 *             ______________________|________________     _____|_________________________________      |
 *            |       Department                      |   |       Company                         |     |
 *            |---------------------------------------|   |---------------------------------------|     |
 *            |-name: String                          |   |-name: String                          |     |
 *            |_______________________________________|   |_______________________________________|<>----
 *            |+Inform(info: String)                  |   |+Inform(info: String)                  |
 *            |+Add(o: Organization)                  |   |+Add(o: Organization)                  |
 *            |+Remove(o: Organization)               |   |+Remove(o: Organization)               |
 *            |+GetName(): string                     |   |+GetName(): string                     |
 *            |+GetChildOrganization(): []Organization|   |+GetChildOrganization(): []Organization|
 *            |_______________________________________|   |_______________________________________|
 *
 *
 *
 * 组合模式（Composite Pattern）将对象组合成树形结构以表示“部分-整体”的层次结构。组合模式使得用户可以使用一致的方法操作单个对象和组合对象。
 *
 * 组合模式角色划分
 * 1. 抽象组件，如上图中的Component
 * 2. 简单组件，如上图中的SimpleComponent
 * 3. 复合组件，如上图中的CompositeComponent
 *
 *
 * 实例介绍
 * 对于一家大型公司，每当公司高层有重要事项需要通知到总部每个部门以及分公司的各个部门时，并不希望逐一通知，而只希望通过总部各部门及分公司，再由分公司通知其所有部门。这样，对于总公司而言，不需要关心通知的是总部的部门还是分公司。
 *
 * 1. 抽象组件:        Organization 抽象组件定义了组件的通知接口，并实现了增删子组件及获取所有子组件的方法。同时重写了hashCode和equales方法
 * 2. 简单组件（部门）: Department   简单组件在通知方法中只负责对接收到消息作出响应。
 * 3. 复合组件（公司）: Company      复合组件在自身对消息作出响应后，还须通知其下所有子组件
 *
 *
 * 组合模式优点
 * 高层模块调用简单。组合模式中，用户不用关心到底是处理简单组件还是复合组件，可以按照统一的接口处理。不必判断组件类型，更不用为不同类型组件分开处理。
 * 组合模式可以很容易的增加新的组件。若要增加一个简单组件或复合组件，只须找到它的父节点即可，非常容易扩展，符合“开放-关闭”原则。
 *
 * 组合模式缺点
 * 无法限制组合组件中的子组件类型。在需要检测组件类型时，不能依靠编译期的类型约束来实现，必须在运行期间动态检测。
 *
 * 已遵循的原则
 * 依赖倒置原则（复合类型不依赖于任何具体的组件而依赖于抽象组件）
 * 迪米特法则
 * 里氏替换原则
 * 接口隔离原则
 * 单一职责原则
 * 开闭原则
 *
 */

// 1. 抽象组件
type Organization interface {
	GetChildOrganization() []Organization
	GetName() string
	Inform(info string)
	Remove(o Organization)
	Add(o Organization)
}

type CommonOrganization struct {
	list []Organization
	name string
}

func (org *CommonOrganization) GetName() string {
	return org.name
}

func (org *CommonOrganization) GetChildOrganization() []Organization {
	return org.list
}

func (org *CommonOrganization) Add(o Organization) {
	org.list = append(org.list, o)
}

func (org *CommonOrganization) Remove(o Organization) {
	for i, v := range org.list {
		if v.GetName() == o.GetName() {
			org.list = append(org.list[:i], org.list[i+1:]...)
		}
	}
}

func (org *CommonOrganization) Inform(info string) {
	println(org.name + " got " + info)
	if nil == org.list {
		return
	}
	for _, v := range org.list {
		println(v.GetName() + " got " + info)
	}
}

// 2. 简单组件
type Department struct {
	CommonOrganization
}

func NewDepartment(name string) *Department {
	return &Department{CommonOrganization{name: name}}
}

// 3. 复合组件
type Company struct {
	CommonOrganization
}

func (cop *Company) Equals(o interface{}) bool { // 重写Equals
	if cop == o {
		return true
	}
	if reflect.TypeOf(cop) == reflect.TypeOf(o) {
		return cop.name == (o.(*Company)).name
	}
	return false
}

func NewCompany(name string) *Company {
	return &Company{CommonOrganization{name: name}}
}
