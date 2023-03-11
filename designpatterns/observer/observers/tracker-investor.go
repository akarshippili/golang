package observers

import "fmt"

type TrackerInvestor struct{}

func (observer *TrackerInvestor) Update(newStockPrice float64) {
	// what the tracker investor wants to do goes here
	fmt.Println("Tracker Inverstor: ", "new stock price", newStockPrice)
}
