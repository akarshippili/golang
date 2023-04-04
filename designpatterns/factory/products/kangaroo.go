package products

import "fmt"

type Kangaroo struct{}

func (kangaroo *Kangaroo) Eat() {
	fmt.Println("kangaroo eating leaves")
}

func (kangaroo *Kangaroo) Talk() {
	fmt.Println("kangaroo making noise")
}

func (kangaroo *Kangaroo) Run() {
	fmt.Println("kangaroo jumping")
}
