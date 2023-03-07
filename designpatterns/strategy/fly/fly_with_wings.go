package fly

import "fmt"

type FlyWithWings struct{}

func (flyBehavior FlyWithWings) Fly() {
	fmt.Println("Flying with wings")
}
