package fly

import "fmt"

type FlyNoWay struct{}

func (flyBehavivor FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
