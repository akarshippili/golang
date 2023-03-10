package main

import (
	"github.com/akarshippili/golang/designpatterns/strategy/duck"
	"github.com/akarshippili/golang/designpatterns/strategy/fly"
	"github.com/akarshippili/golang/designpatterns/strategy/quack"
)

func main() {

	duck := duck.Duck{
		FlyBehavior:   fly.FlyWithWings{},
		QuackBehavior: quack.QuackingBehavior{},
	}
	duck.Behave()

	duck.SetFlyingBehavior(fly.FlyNoWay{})
	duck.SetQuackingBehavior(quack.MuteBehavior{})
	duck.Behave()

	duck.SetFlyingBehavior(fly.FlyRocketPowred{})
	duck.SetQuackingBehavior(quack.SqueakBehavior{})
	duck.Behave()
}
