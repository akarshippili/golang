package quack

import "fmt"

type MuteBehavior struct{}

func (quack MuteBehavior) Quack() {
	fmt.Println("I can't quack")
}
