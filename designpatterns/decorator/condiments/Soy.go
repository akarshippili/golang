package condiments

import "github.com/akarshippili/golang/designpatterns/decorator/beverages"

type Soy struct {
	Beverage beverages.Beverage
}

func (b Soy) GetCost() float64 {
	return b.Beverage.GetCost() + 0.15
}

func (b Soy) GetDescription() string {
	return b.Beverage.GetDescription() + ", " + "Soy"
}
