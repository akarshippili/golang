package condiments

import "github.com/akarshippili/golang/designpatterns/decorator/beverages"

type Whip struct {
	Beverage beverages.Beverage
}

func (b Whip) GetCost() float64 {
	return b.Beverage.GetCost() + 0.15
}

func (b Whip) GetDescription() string {
	return b.Beverage.GetDescription() + ", " + "Whip"
}
