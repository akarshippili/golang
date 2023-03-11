package main

import (
	"github.com/akarshippili/golang/designpatterns/observer/observers"
	"github.com/akarshippili/golang/designpatterns/observer/subject"
)

func main() {
	stockTicker := subject.InitStockTicker(0.0)

	buyerInvestor := observers.BuyerInvestor{}
	sellerInvestor := observers.SellerInvestor{}
	trackerInvestor := observers.TrackerInvestor{}

	stockTicker.Register(&buyerInvestor)
	stockTicker.Register(&sellerInvestor)
	stockTicker.Register(&trackerInvestor)

	stockTicker.SetStockPrice(1.1)
	stockTicker.Remove(&sellerInvestor)

	stockTicker.SetStockPrice(1.2)
}
