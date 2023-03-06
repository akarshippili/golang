package fly

import "fmt"

type FlyWithWings struct{}

func (flyBehavivor FlyWithWings) Fly() {
	fmt.Println("Flying with wings")
}
