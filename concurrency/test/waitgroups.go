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
}
