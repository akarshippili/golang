package beverages

type HouseBlend struct{}

func (bevrage HouseBlend) GetCost() float64 {
	return 0.89
}

func (bevrage HouseBlend) GetDescription() string {
	return "House Blend"
}
