package _4_decorator

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
