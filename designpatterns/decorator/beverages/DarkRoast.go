package beverages

type DarkRoast struct{}

func (bevrage DarkRoast) GetCost() float64 {
	return 0.99
}

func (bevrage DarkRoast) GetDescription() string {
	return "Dark Roast"
}
