package beverages

type Espresso struct{}

func (bevrage Espresso) GetCost() float64 {
	return 1.99
}

func (bevrage Espresso) GetDescription() string {
	return "Espresso"
}
