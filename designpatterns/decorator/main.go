package main

import (
	"fmt"

	"github.com/akarshippili/golang/designpatterns/decorator/beverages"
	"github.com/akarshippili/golang/designpatterns/decorator/condiments"
)

func main() {
	var drink beverages.Beverage = beverages.Decaf{}
	drink = condiments.Whip{
		Beverage: drink,
	}
	drink = condiments.Whip{
		Beverage: drink,
	}
	drink = condiments.Soy{
		Beverage: drink,
	}
	drink = condiments.Mocha{
		Beverage: drink,
	}

	fmt.Println(drink.GetCost(), drink.GetDescription())
}
