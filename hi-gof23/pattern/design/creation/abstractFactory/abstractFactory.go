package abstractFactory

/**
 * 在工厂方法模式中一种工厂只能创建一种具体产品。而在抽象工厂模式中一种具体工厂可以创建多个种类的具体产品。
 *
 *
 * 抽象工厂模式介绍
 * 抽象工厂模式（Abstract Factory Method Pattern）中，抽象工厂提供一系列创建多个抽象产品的接口，而具体的工厂负责实现具体的产品实例。抽象工厂模式与工厂方法模式最大的区别在于抽象工厂中每个工厂可以创建多个种类的产品。
 *
 *
 * 抽象工厂模式角色划分
 * 1. 抽象产品（或者产品接口），如上文类图中的IUserDao，IRoleDao，IProductDao
 * 2. 具体产品，如PostgreSQLProductDao
 * 3. 抽象工厂（或者工厂接口），如IFactory
 * 4. 具体工厂，如果MySQLFactory
 * 5. 产品族，如Oracle产品族，包含OracleUserDao，OracleRoleDao，OracleProductDao
 *
 *
 * 抽象工厂模式角色划分
 * 1. 抽象产品（或者产品接口） 如                  Car
 * 2. 具体产品，            如                  BenzCar, BMWCar, LandRoverBusCar, LandRoverMiniCar, LandRoverSUVCar
 * 3. 产品族，              如(LandRover产品族)  LandRoverBusCar, LandRoverMiniCar, LandRoverSUVCar
 *
 * 4. 抽象工厂（或者工厂接口），   如 CarFamilyFactory
 * 5. 具体工厂，                如 LandRoverCarFamilyFactory
 *
 *
 * 抽象工厂模式优点
 * 1. 因为每个具体工厂类只负责创建产品，没有简单工厂中的逻辑判断，因此符合单一职责原则。
 * 2. 与简单工厂模式不同，抽象工厂并不使用静态工厂方法，可以形成基于继承的等级结构。
 * 3. 新增一个产品族（如上文类图中的MySQLUserDao，MySQLRoleDao，MySQLProductDao）时，只需要增加相应的具体产品和对应的具体工厂类即可。相比于简单工厂模式需要修改判断逻辑而言，抽象工厂模式更符合开-闭原则。
 *
 * 抽象工厂模式缺点
 * 1. 新增产品种类（如上文类图中的UserDao，RoleDao，ProductDao）时，需要修改工厂接口（或者抽象工厂）及所有具体工厂，此时不符合开-闭原则。抽象工厂模式对于新的产品族符合开-闭原则而对于新的产品种类不符合开-闭原则，这一特性也被称为开-闭原则的倾斜性。
 *
 * 已遵循的原则
 * 1. 依赖倒置原则（工厂构建产品的方法均返回产品接口而非具体产品，从而使客户端依赖于产品抽象而非具体）
 * 2. 迪米特法则
 * 3. 里氏替换原则
 * 4. 接口隔离原则
 * 5. 单一职责原则（每个工厂只负责创建自己的具体产品族，没有简单工厂中的逻辑判断）
 * 6. 开闭原则（增加新的产品族，不像简单工厂那样需要修改已有的工厂，而只需增加相应的具体工厂类）
 *
 * 未遵循的原则
 * 1. 开闭原则（虽然对新增产品族符合开-闭原则，但对新增产品种类不符合开-闭原则）
 *
 **/

/**
 * 在工厂方法模式中具体工厂负责生产具体的产品，每一个具体工厂对应一种具体产品，工厂方法也具有唯一性，一般情况下，一个具体工厂中只有一个工厂方法或者一组重载的工厂方法。但是有时候我们需要一个工厂可以提供多个产品对象，而不是单一的产品对象。
 *
 * 为了更清晰地理解工厂方法模式，需要先引入两个概念：
 * 产品等级结构 ：产品等级结构即产品的继承结构，如一个抽象类是电视机，其子类有海尔电视机、海信电视机、TCL电视机，则抽象电视机与具体品牌的电视机之间构成了一个产品等级结构，抽象电视机是父类，而具体品牌的电视机是其子类。
 * 产品族 ：在抽象工厂模式中，产品族是指由同一个工厂生产的，位于不同产品等级结构中的一组产品，如海尔电器工厂生产的海尔电视机、海尔电冰箱，海尔电视机位于电视机产品等级结构中，海尔电冰箱位于电冰箱产品等级结构中。当系统所提供的工厂所需生产的具体产品并不是一个简单的对象，而是多个位于不同产品等级结构中属于不同类型的具体产品时需要使用抽象工厂模式。
 *
 * 抽象工厂模式是所有形式的工厂模式中最为抽象和最具一般性的一种形态。
 * 抽象工厂模式与工厂方法模式最大的区别在于，工厂方法模式针对的是一个产品等级结构，而抽象工厂模式则需要面对多个产品等级结构，一个工厂等级结构可以负责多个不同产品等级结构中的产品对象的创建 。当一个工厂等级结构可以创建出分属于不同产品等级结构的一个产品族中的所有对象时，抽象工厂模式比工厂方法模式更为简单、有效率。
 * 抽象工厂模式(Abstract Factory Pattern)：提供一个创建一系列相关或相互依赖对象的接口，而无须指定它们具体的类。抽象工厂模式又称为Kit模式，属于对象创建型模式。
 *
 * 抽象工厂模式包含如下角色：
 * 1. AbstractFactory：抽象工厂
 * 2. ConcreteFactory：具体工厂
 * 3. AbstractProduct：抽象产品
 * 4. Product：具体产品
 *
 */

type PlantFactory interface {
	MakePlant() Plant
	MakePicker() Picker
}

type Plant interface {
}

type Picker interface {
}

type PearFactory struct {
}

type Pear struct {
	desc string
}

type PearPicker struct {
	desc string
}

func (_ PearFactory) MakePlant() Plant {
	return Pear{"我是一个梨"}
}

func (_ PearFactory) MakePicker() Picker {
	return PearPicker{"采摘一个梨"}
}
