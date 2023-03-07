package duck

import (
	"github.com/akarshippili/golang/designpatterns/strategy/fly"
	"github.com/akarshippili/golang/designpatterns/strategy/quack"
)

type Duck struct {
	Fly   fly.Fly
	Quack quack.Quack
}

func (duck *Duck) SetFlyingBehavior(fly fly.Fly) {
	duck.Fly = fly
}

func (duck *Duck) SetQuackingBehavior(quack quack.Quack) {
	duck.Quack = quack
}

func (duck *Duck) Behave() {
	duck.Quack.Quack()
	duck.Quack.Quack()
	duck.Fly.Fly()
}
