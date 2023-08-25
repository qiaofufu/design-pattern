package _6_factory_method

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
