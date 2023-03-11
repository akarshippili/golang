package observers

import "fmt"

type BuyerInvestor struct{}

func (observer *BuyerInvestor) Update(newStockPrice float64) {
	// what the buyer investor wants to do goes here
	fmt.Println("Buyer Inverstor: ", "new stock price", newStockPrice)
}
