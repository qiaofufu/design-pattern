package _5_proxy

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
