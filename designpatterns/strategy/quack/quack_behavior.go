package quack

import "fmt"

type QuackingBehavior struct{}

func (quack QuackingBehavior) Quack() {
	fmt.Println("Quack")
}
