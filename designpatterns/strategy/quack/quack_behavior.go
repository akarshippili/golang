package quack

import "fmt"

type QuackBehavior struct{}

func (quack QuackBehavior) Quack() {
	fmt.Println("Quack")
}
