package bridge

/**
 * 不必要的继承导致类爆炸
 * 汽车可按品牌分（本例中只考虑BMT，BenZ），也可按手动档、自动档、手自一体来分。
 * 如果对于每一种车都实现一个具体类，则一共要实现3*3=9个类。
 *
 *
 *
 * 桥接模式类图
 *
 *               ___________________                     ________________________
 *              |   AbstractCar     |                   |       Transmission     |
 *              |-------------------|                   |------------------------|
 *              |-strategy:Strategy |<>---------------->|+strategy(String input) |
 *              |-------------------|                   |------------------------|
 *              |+run(): void       |                   |+gear(): void           |
 *              |___________________|                   |________________________|
 *                       /\                                               /\
 *                       |                                                |
 *        _______________|____________                     _______________|_______
 *   ____|______                 ____|______         _____|________          _____|________
 *  |  BMWCar   |               |  BenZCar  |       | Auto         |        | Manual       |
 *  |-----------|               |-----------|       |--------------|        |--------------|
 *  |           |               |           |       |              |        |              |
 *  |-----------|               |-----------|       |--------------|        |--------------|
 *  |+run():void|               |+run():void|       |+gear(): void |        |+gear(): void |
 *  |___________|               |___________|       |______________|        |______________|
 *
 *
 * 桥接模式（Bridge Pattern），将抽象部分与它的实现部分分离，使它们都可以独立地变化。
 * 更容易理解的表述是：实现系统可从多种维度分类，桥接模式将各维度抽象出来，各维度独立变化，之后可通过聚合，将各维度组合起来，减少了各维度间的耦合。
 *
 * // 从上图可以看到，对于每种组合都需要创建一个具体类，如果有N个维度，每个维度有M种变化，则需要MN个具体类，类非常多，并且非常多的重复功能。
 * // 如果某一维度，如Transmission多一种可能，比如手自一体档（AMT），则需要增加3个类，BMWAMT，BenZAMT，LandRoverAMT。
 *
 * 从上图可知，当把每个维度拆分开来，只需要M*N个类，并且由于每个维度独立变化，基本不会出现重复代码。
 * 此时如果增加手自一体档，只需要增加一个AMT类即可
 *
 *
 * 桥接模式意图: 桥接模式将抽象部分与它的实现部分分离，使它们可以独立变化
 *
 *
 *
 * 桥接模式角色划分:
 * 1. 抽象类接口
 * 2. 分离的抽象类接口
 * 3. 抽象类接口实现
 * 4. 分离的抽象类接口实现
 *
 *
 * 桥接模式适用场景:
 * 1. 不希望在抽象和它的实现部分之间有一个固定的绑定关系。比如这种情况可能是因为在程序运行时刻实现部分应可以被选择或者切换。
 * 2. 类的抽象以及它的实现应该可以通过生成子类的方法加以扩充。
 * 3. 对一个抽象的实现部分的修改应对客户不产生影响，即客户的代码不必重新编译。
 * 4. 一个类型存在两个独立变化的维度，且这两个维度都需要进行扩展。
 *
 *
 * 桥接模式优点:
 * 1. 分离接口及其实现部分：一个实现未必不变地绑定在一个接口上，抽象类的实现可在运行时刻进行配置，一个对象甚至可以在运行时刻改变它的实现。将abstraction与Implementor分离有助于降低对实现部分编译时刻的依赖性，当改变一个实现类时，不需要重新编译abstraction类和客户重新。为了保证一个类库的不同版本之间二进制兼容性，一定要有这个性质。另外，接口和实现分离有助于分层，从而产生更好的结构化系统，系统的高层部分只要知道abstraction和implementor即可。
 * 2. 提高可扩展性，可以独立对Abstraction和Implementor层次进行扩展。
 * 3. 实现细节对可对客户透明。
 *
 * 桥接模式缺点:
 * 1. 不容易设计，需不需要分离，如何分离等问题。比较难以拿捏。
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
 */

// 1. 抽象类接口
type AbstractCar struct {
	trn Transmission
}

func (a *AbstractCar) SetTransmission(trn Transmission) { // *AbstractCar 必须使用指针, 子类继承时需要指定内部元素
	a.trn = trn
}

// 2. 分离的抽象类接口
type Transmission interface {
	Gear()
}

// 3. 抽象类接口实现
type BenZCar struct {
	AbstractCar
}

func (car BenZCar) run() {
	car.AbstractCar.trn.Gear()
	println("BenZCar is running")
}

type BMWCar struct {
	AbstractCar
}

func (car BMWCar) run() {
	car.AbstractCar.trn.Gear()
	println("BMWCar is running")
}

// 4. 分离的抽象类接口实现
type Manual struct { // 手动档
}

func (trn Manual) Gear() {
	println("Manual transmission")
}

type Auto struct { // 自动档
}

func (trn Auto) Gear() {
	println("Auto transmission")
}

type AMT struct { // 手自一体档
}

func (trn AMT) Gear() {
	println("AMT transmission")
}