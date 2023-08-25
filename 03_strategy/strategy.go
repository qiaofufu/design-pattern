package _3_strategy

type Result struct {
	Instance []string
	CacheKey string
}

type Balancer struct {
	strategy BalanceStrategy
}

func New(strategy BalanceStrategy) Balancer {
	return Balancer{
		strategy: strategy,
	}
}

func (b *Balancer) Peek() Result {
	return b.strategy.Peek()
}

func (b *Balancer) Delete() {
	b.strategy.Delete()
}
func (b *Balancer) ReBalancing() {
	b.strategy.ReBalancing()
}

type BalanceStrategy interface {
	Peek() Result
	Delete()
	ReBalancing()
}

type RoundRobin struct {
}

func (r RoundRobin) Peek() Result {
	//TODO implement me
	panic("implement me")
}

func (r RoundRobin) Delete() {
	//TODO implement me
	panic("implement me")
}

func (r RoundRobin) ReBalancing() {
	//TODO implement me
	panic("implement me")
}

type WeightRoundRobin struct {
}

func (w WeightRoundRobin) Peek() Result {
	//TODO implement me
	panic("implement me")
}

func (w WeightRoundRobin) Delete() {
	//TODO implement me
	panic("implement me")
}

func (w WeightRoundRobin) ReBalancing() {
	//TODO implement me
	panic("implement me")
}
