package _2_facade

import "testing"

func TestFacade(t *testing.T) {
	facade := NewFacade()
	if facade.Test() != "model1"+"model2" {
		t.Fatalf("failed to impl facade")
	}
}
