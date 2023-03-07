package fly

import "fmt"

type FlyRocketPowred struct{}

func (flyBehavior FlyRocketPowred) Fly() {
	fmt.Println("Flying With Rocket Powered Wings")
}
