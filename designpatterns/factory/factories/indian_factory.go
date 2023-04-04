package factories

import (
	"math/rand"

	"github.com/akarshippili/golang/designpatterns/factory/products"
)

type IndianFactory struct{}

func (f *IndianFactory) GetAnimal() products.Animal {

	random := rand.Intn(4)

	switch random {
	case 0:
		return new(products.BengalTiger)
	case 1:
		return new(products.IndianLion)
	case 2:
		return new(products.KingCobra)
	case 3:
		return new(products.Dog)
	default:
		return nil
	}

}
