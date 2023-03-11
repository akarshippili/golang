package observers

// push model
type Investor interface {
	Update(newStockPrice float64)
}
