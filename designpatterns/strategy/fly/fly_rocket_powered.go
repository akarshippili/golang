package fly

import "fmt"

type FlyRocketPowred struct{}

func (flyBehavivor FlyRocketPowred) Fly() {
	fmt.Println("Flying With Rocket Powered Wings")
}
