package observers

import "fmt"

type TrackerInvestor struct{}

func (observer *TrackerInvestor) Update(newStockPrice float64) {
	// what the tracker investor wants to do with the new price goes here
	fmt.Println("Tracker Inverstor: ", "new stock price", newStockPrice)
}
