package duck

import (
	"github.com/akarshippili/golang/designpatterns/strategy/fly"
	"github.com/akarshippili/golang/designpatterns/strategy/quack"
)

type Duck struct {
	FlyBehavior   fly.FlyBehavior
	QuackBehavior quack.QuackBehavior
}

func (duck *Duck) SetFlyingBehavior(fly fly.FlyBehavior) {
	duck.FlyBehavior = fly
}

func (duck *Duck) SetQuackingBehavior(quack quack.QuackBehavior) {
	duck.QuackBehavior = quack
}

func (duck *Duck) Behave() {
	duck.QuackBehavior.Quack()
	duck.QuackBehavior.Quack()
	duck.FlyBehavior.Fly()
}
