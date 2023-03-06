package duck

import (
	"github.com/akarshippili/golang/designpatterns/strategy/fly"
	"github.com/akarshippili/golang/designpatterns/strategy/quack"
)

type Duck struct {
	fly   fly.Fly
	quack quack.Quack
}

func (duck *Duck) setFlyingBehavior(fly fly.Fly) {
	duck.fly = fly
}

func (duck *Duck) setQuackingBehavior(fly fly.Fly) {
	duck.fly = fly
}
