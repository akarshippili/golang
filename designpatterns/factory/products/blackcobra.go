package products

import "fmt"

type BlackCobra struct{}

func (blackCobra *BlackCobra) Eat() {
	fmt.Println("blackCobra eating rats")

}
func (blackCobra *BlackCobra) Talk() {
	fmt.Println("blackCobra hissing")
}
func (blackCobra *BlackCobra) Run() {
	fmt.Println("blackCobra slithering")
}
