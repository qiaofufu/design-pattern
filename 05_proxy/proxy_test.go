package _5_proxy

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
