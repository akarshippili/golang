package products

import "fmt"

type KingCobra struct{}

func (kingCobra *KingCobra) Eat() {
	fmt.Println("kingCobra eating rats")

}
func (kingCobra *KingCobra) Talk() {
	fmt.Println("kingCobra hissing")
}
func (kingCobra *KingCobra) Run() {
	fmt.Println("kingCobra slithering")
}
