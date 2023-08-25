package _1_simple_factory

import "testing"

func TestNewAPI(t *testing.T) {
	api := New(1)
	if api.Format("cill") != "[hello] cill" {
		t.Fatalf("failed to impl API interface{}")
	}
	api2 := New(2)
	if api2.Format("cill") != "[hi] cill" {
		t.Fatalf("failed to impl API interface{}")
	}
}
