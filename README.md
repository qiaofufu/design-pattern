# 设计模式

## 创建型模式

### 简单工厂模式

#### 应用场景

> 当有一个接口有多种实现时, 我们根据需要选择创建不同的对象时, 可以使用简单工厂, 不需要向外部暴露接口的具体实现, 让用户对接口的具体实现无感

#### 代码实现

```go
type API interface {
	Format(msg string) string
}

type testAPI1 struct { }

func (t testAPI1) Format(msg string) string {
	return fmt.Sprint("[hello] " + msg)
}

type testAPI2 struct { }

func (t testAPI2) Format(msg string) string {
	return fmt.Sprintf("[hi] " + msg)
}

func New(typeId int) API {
	switch typeId {
	case 1:
		return testAPI1{}
	case 2:
		return testAPI2{}
	default:
		return testAPI1{}
	}
}
```

### 工厂方法模式

#### 应用场景

> 简单工厂模式违背了开放封闭原则,  根据依赖倒转原则, 我们可以通过添加工厂接口,  让工厂接口的子类去生产对应的对象, 将创建对象时的逻辑判断移动到客户端去实现(简单工程每次添加新的实现,需要添加case分支)

#### 代码实现

```go factory_method.go
type API interface {
	Call() string
}

type IFactory interface {
	Create() API
}

type helloAPI struct{}

func (*helloAPI) Call() string {
	return "hello"
}

type HelloAPIFactory struct{}

func (HelloAPIFactory) Create() API {
	return &helloAPI{}
}

type hiAPI struct{}

func (*hiAPI) Call() string {
	return "hi"
}

type HiAPIFactory struct{}

func (HiAPIFactory) Create() API {
	return &hiAPI{}
}
```

```go factory_method_test.go
import (
	"reflect"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	helloAPI := HelloAPIFactory{}.Create()
	if !reflect.DeepEqual(helloAPI.Call(), "hello") {
		t.Fatalf("helloapi failed to impl factory method")
	}
	hiAPI := HiAPIFactory{}.Create()
	if !reflect.DeepEqual(hiAPI.Call(), "hi") {
		t.Fatalf("hiapi failed to impl factory method")
	}
}
```

## 结构型模式

### 外观模式

#### 应用场景

> 当我们需要一个系统中的多个子系统提供的功能时, 我们可以将子系统的功能集成起来, 仅向外暴露少部分集成后的功能, 降低使用者的学习成本, 比如当你使用http.Get() 函数时, 你并不需要去了解它内部http报文拼装和发送请求的代码调用, 你只需执行一个函数, 不需要了解他具体是如何执行的

#### 代码实现

```go
type API interface {
	Test() string
}

func NewFacade() API {
	return apiImpl{}
}

type apiImpl struct { }

// 外部通过调用Test()即可完成操作, 不需要自己再去了解Model1 和 Model2的用法
func (a apiImpl) Test() string { 
	model1 := Model1{}.TestModel1()
	model2 := Model2{}.TestModel2()
	return model1 + model2
}

type Model1 struct { }

func (m Model1) TestModel1() string {
	return "model1"
}

type Model2 struct { }

func (m Model2) TestModel2() string {
	return "model2"
}

```

### 装饰模式

#### 应用场景

> 当我们想在运行时给一个对象动态的添加一系列职责(功能)(按顺序),又不会更改原本对象时. 比如: 算加减法,按照从前往后的方式1+2=3, 1+2+3=(1+2) + 3

#### 代码实现

```go decorator.go
type Component interface {
	Call() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Call() int {
	return 0
}

type AddDecorator struct {
	component Component
	num       int
}

func WarpAddDecorator(component Component, num int) Component {
	return &AddDecorator{component: component, num: num}
}

func (a AddDecorator) Call() int {
	return a.component.Call() + a.num
}

type SubDecorator struct {
	component Component
	num       int
}

func WarpSubDecorator(component Component, num int) Component {
	return &SubDecorator{
		component: component,
		num:       num,
	}
}

func (s *SubDecorator) Call() int {
	return s.component.Call() + s.num	
}
```
```go decorator_test.go
import (
	"reflect"
	"testing"
)

func TestDecorator(t *testing.T) {
	var component Component
	component = new(ConcreteComponent)
	component = WarpAddDecorator(component, 1)
	component = WarpSubDecorator(component, 2)
	if !reflect.DeepEqual(component.Call(), 3) {
		t.Fatalf("failed to impl Decorator pattern")
	}
}
```

### 代理模式

#### 应用场景

> 顾名思义, 就是代理去做某事, 可能远程访问/访问控制/记录日志/缓存....

#### 代码实现

```go proxy.go
type Subject interface {
	Do() string
}

type RealSubject struct{}

func (r RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var result string
	// Do 之前进行处理
	result += "pre:"
	result += p.real.Do()
	// Do 之后进行处理
	result += ":after"
	return result
}
```
```go proxy_test.go
import (
	"reflect"
	"testing"
)

func TestProxy(t *testing.T) {
	res := Proxy{}.Do()
	if !reflect.DeepEqual(res, "pre:real:after") {
		t.Fatalf("failed to proxy")
	}
}

```


## 行为型模式

### 策略模式

#### 应用场景

> 当的操作(算法/策略)需要动态的变化(运行时更改)时, 可以将操作(算法/策略)分离出来

#### 代码实现

```go
type Result struct {
	Instance []string
	CacheKey string
}

// 通过Balancer对象封装Strategy, 其中的strategy可以在运行时变化
type Balancer struct {
	strategy BalanceStrategy
}

func New(strategy BalanceStrategy) Balancer {
	return Balancer{
		strategy: strategy,
	}
}

func (b *Balancer) Peek() Result {
	return b.strategy.Peek()
}

func (b *Balancer) Delete() {
	b.strategy.Delete()
}

func (b *Balancer) ReBalancing() {
	b.strategy.ReBalancing()
}

type BalanceStrategy interface {
	Peek() Result
	Delete()
	ReBalancing()
}

type RoundRobin struct { }

func (r RoundRobin) Peek() Result {
	//TODO implement me
	panic("implement me")
}

func (r RoundRobin) Delete() {
	//TODO implement me
	panic("implement me")
}

func (r RoundRobin) ReBalancing() {
	//TODO implement me
	panic("implement me")
}

type WeightRoundRobin struct { }

func (w WeightRoundRobin) Peek() Result {
	//TODO implement me
	panic("implement me")
}

func (w WeightRoundRobin) Delete() {
	//TODO implement me
	panic("implement me")
}

func (w WeightRoundRobin) ReBalancing() {
	//TODO implement me
	panic("implement me")
}
```