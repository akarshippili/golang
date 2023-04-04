package products

import "fmt"

type Dog struct{}

func (dag *Dog) Eat() {
	fmt.Println("dog eating food")
}
func (dog *Dog) Talk() {
	fmt.Println("dog barking")
}
func (dog *Dog) Run() {
	fmt.Println("dog running")
}
