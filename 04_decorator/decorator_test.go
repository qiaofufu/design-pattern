package _4_decorator

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
