package factories

import (
	"math/rand"

	"github.com/akarshippili/golang/designpatterns/factory/products"
)

type AfricanFactory struct{}

func (f *AfricanFactory) GetAnimal() products.Animal {

	random := rand.Intn(4)

	switch random {
	case 0:
		return new(products.AfricanLion)
	case 1:
		return new(products.AfricanTiger)
	case 2:
		return new(products.BlackCobra)
	case 3:
		return new(products.Dog)
	default:
		return nil
	}

}
