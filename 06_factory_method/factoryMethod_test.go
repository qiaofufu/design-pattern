package _6_factory_method

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
