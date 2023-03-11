package subject

import "github.com/akarshippili/golang/designpatterns/observer/observers"

type StockTicker struct {
	stockPrice float64
	investors  map[observers.Investor]bool
}

func InitStockTicker(currentStockPrice float64) StockTicker {
	stockTicker := StockTicker{
		stockPrice: currentStockPrice,
		investors:  make(map[observers.Investor]bool),
	}

	return stockTicker
}

func (stockTicker *StockTicker) GetStockPrice() float64 {
	return stockTicker.stockPrice
}

func (stockTicker *StockTicker) SetStockPrice(newStockPrice float64) {
	stockTicker.stockPrice = newStockPrice
	stockTicker.Notify()
}

func (stockTicker *StockTicker) Register(observer observers.Investor) {
	if _, ok := stockTicker.investors[observer]; !ok {
		stockTicker.investors[observer] = true
	}
}

func (stockTicker *StockTicker) Remove(observer observers.Investor) {
	delete(stockTicker.investors, observer)
}

func (stockTicker *StockTicker) Notify() {
	for observer := range stockTicker.investors {
		observer.Update(stockTicker.GetStockPrice())
	}
}
