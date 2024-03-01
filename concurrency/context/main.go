package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func(ctx context.Context) {
	// 	defer wg.Done()
	// 	for val := range time.Tick(time.Second) {
	// 		if ctx.Err() != nil {
	// 			fmt.Println(ctx.Err().Error())
	// 			return
	// 		}
	// 		fmt.Printf("Tick... %v \n", val)
	// 	}
	// }(ctx)

	// time.Sleep(4 * time.Second)
	// cancel()

	// wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for val := range time.Tick(time.Second) {
			if ctx.Err() != nil {
				fmt.Println(ctx.Err().Error())
				return
			}
			fmt.Printf("Tick... %v \n", val)
		}
	}(ctx)

	time.Sleep(4 * time.Second)

	wg.Wait()
}
