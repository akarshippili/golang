package fly

import "fmt"

type FlyNoWay struct{}

func (flyBehavior FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
