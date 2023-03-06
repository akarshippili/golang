package quack

import "fmt"

type SqueakBehavior struct{}

func (quack SqueakBehavior) Quack() {
	fmt.Println("Squeak squeak!!!")
}
