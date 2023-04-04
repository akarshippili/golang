package products

import "fmt"

type AfricanLion struct{}

func (africanLion *AfricanLion) Eat() {
	fmt.Println("africanLion eating meat")
}
func (africanLion *AfricanLion) Talk() {
	fmt.Println("africanLion roaring")
}

func (africanLion *AfricanLion) Run() {
	fmt.Println("africanLion running")
}
