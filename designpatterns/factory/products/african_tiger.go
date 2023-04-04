package products

import "fmt"

type AfricanTiger struct{}

func (tiger *AfricanTiger) Eat() {
	fmt.Println("AfricanTiger eating meat")
}

func (tiger *AfricanTiger) Talk() {
	fmt.Println("AfricanTiger roaring")
}

func (tiger *AfricanTiger) Run() {
	fmt.Println("AfricanTiger running")
}
