package factories

import (
	"math/rand"

	"github.com/akarshippili/golang/designpatterns/factory/products"
)

type RandomFactory struct{}

func (f *RandomFactory) GetAnimal() products.Animal {

	random := rand.Intn(8)

	switch random {
	case 0:
		return new(products.AfricanLion)
	case 1:
		return new(products.AfricanTiger)
	case 2:
		return new(products.BengalTiger)
	case 3:
		return new(products.BlackCobra)
	case 4:
		return new(products.IndianLion)
	case 5:
		return new(products.Kangaroo)
	case 6:
		return new(products.KingCobra)
	case 7:
		return new(products.Dog)
	default:
		return nil
	}

}
