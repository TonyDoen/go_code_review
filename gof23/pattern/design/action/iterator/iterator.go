package iterator

/**
 * 迭代器模式是一种设计模式，是一种最简单也最常见的设计模式。
 * 它可以让使用者透过特定的介面巡访容器中的每一个元素而不用了解底层的实作。(迭代子模式可以顺序地访问一个聚集中的元素而不必暴漏聚集的内部表象)
 * 简单点说，为一个容器设置一个迭代函数，可以使用这个迭代函数来顺序访问其中的每一个元素，而外部无需知道底层实现。
 * 如果再结合 访问者模式,向其中传入自定义的访问者，那么就可以让访问者访问容器中的每个元素了
 *
 *
 * 迭代器模式角色
 * 1. 抽象聚合类: 定义一个抽象的容器
 * 2. 具体聚合类: 实现上面的抽象类，作为一个容器，用来存放元素，等待迭代
 * 3. 抽象迭代器: 迭代器接口，每个容器下都有一个该迭代器接口的具体实现
 * 4. 具体迭代器: 根据不同的容器，需要定义不同的具体迭代器，定义了游标移动的具体实现
 *
 *
 *
 * 迭代器模式优点
 * 1. 迭代子模式简化了聚集的接口。迭代子具备了一个遍历接口，这样聚集的接口就不必具备遍历接口。
 * 2. 每一个聚集对象都可以有一个或多个迭代子对象，每一个迭代子的迭代状态可以是彼此独立的。因此，一个聚集对象可以同时有几个迭代在进行之中。
 * 3. 由于遍历算法被封装在迭代子角色里面，因此迭代的算法可以独立于聚集角色变化。
 *
 *
 *
 *
 *
 *
 */

// 1. 创建抽象容器
type Aggregate interface {
	Iterator() Iterator
}

// 2. 创建抽象迭代器
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// 3. 创建具体容器
type ConcreteAggregate struct {
	container []interface{} // 容器中装载 interface{} 型容器
}

// 创建一个迭代器，并让迭代器中的容器指针指向当前对象
func (ca *ConcreteAggregate) Iterator() Iterator {
	i := &ConcreteIterator{size: len(ca.container), cursor: 0, aggregate: ca}
	return i
}

// 4. 创建具体迭代器
type ConcreteIterator struct {
	size      int
	cursor    int                // 当前游标
	aggregate *ConcreteAggregate // 对应的容器指针
}

// 判断是否迭代到最后，如果没有，则返回true
func (i *ConcreteIterator) HasNext() bool {
	if i.cursor < i.size {
		return true
	}
	return false
}

// 获取当前迭代元素, 并将游标指向下一个元素
func (i *ConcreteIterator) Next() interface{} {
	if i.cursor < i.size {
		res := i.aggregate.container[i.cursor]
		i.cursor += 1
		return res
	}
	return nil
}
