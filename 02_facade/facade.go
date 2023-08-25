package _2_facade

type API interface {
	Test() string
}

func NewFacade() API {
	return apiImpl{}
}

type apiImpl struct {
}

func (a apiImpl) Test() string {
	model1 := Model1{}.TestModel1()
	model2 := Model2{}.TestModel2()
	return model1 + model2
}

type Model1 struct {
}

func (m Model1) TestModel1() string {
	return "model1"
}

type Model2 struct {
}

func (m Model2) TestModel2() string {
	return "model2"
}
