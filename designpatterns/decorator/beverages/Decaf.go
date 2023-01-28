package beverages

type Decaf struct{}

func (bevrage Decaf) GetCost() float64 {
	return 1.05
}

func (bevrage Decaf) GetDescription() string {
	return "Decaf"
}
