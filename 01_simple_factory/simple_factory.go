package _1_simple_factory

import "fmt"

type API interface {
	Format(msg string) string
}

type testAPI1 struct {
}

func (t testAPI1) Format(msg string) string {
	return fmt.Sprint("[hello] " + msg)
}

type testAPI2 struct {
}

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
