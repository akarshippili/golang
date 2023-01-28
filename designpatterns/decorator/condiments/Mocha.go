package condiments

import "github.com/akarshippili/golang/designpatterns/decorator/beverages"

type Mocha struct {
	Beverage beverages.Beverage
}

func (b Mocha) GetCost() float64 {
	return b.Beverage.GetCost() + 0.15
}

func (b Mocha) GetDescription() string {
	return b.Beverage.GetDescription() + ", " + "Mocha"
}
