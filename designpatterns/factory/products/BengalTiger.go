package products

import "fmt"

type BengalTiger struct{}

func (tiger *BengalTiger) Eat() {
	fmt.Println("BengalTiger eating meat")
}

func (tiger *BengalTiger) Talk() {
	fmt.Println("BengalTiger roaring")
}

func (tiger *BengalTiger) Run() {
	fmt.Println("BengalTiger running")
}
