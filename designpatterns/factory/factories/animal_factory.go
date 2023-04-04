package factories

import "github.com/akarshippili/golang/designpatterns/factory/products"

type AnimalFactory interface {
	GetAnimal() products.Animal
}
