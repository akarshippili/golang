package observers

import "fmt"

type SellerInvestor struct{}

func (observer *SellerInvestor) Update(newStockPrice float64) {
	// what the seller investor wants to do with the new price goes here
	fmt.Println("Seller Inverstor: ", "new stock price", newStockPrice)
}
