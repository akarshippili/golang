package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("async thing")
		wg.Done()
	}()

	wg.Wait()

	ChannelTest()
	SelectChannelTest()
	LoopThroughChannel()
}

// dealdlock reads and writes are blocked util corosponding operation is ready
// func main() {
// 	ch := make(chan int)
// 	ch <- 1
// }
