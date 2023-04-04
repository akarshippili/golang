package products

import "fmt"

type IndianLion struct{}

func (lion *IndianLion) Eat() {
	fmt.Println("IndianLion eating meat")
}
func (lion *IndianLion) Talk() {
	fmt.Println("IndianLion roaring")
}

func (lion *IndianLion) Run() {
	fmt.Println("IndianLion running")
}
