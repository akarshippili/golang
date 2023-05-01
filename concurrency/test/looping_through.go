package main

import (
	"fmt"
	"time"
)

func LoopThroughChannel() {

	ch := make(chan string)

	go func() {
		fmt.Println("Writing Hello1 to channel")
		ch <- "Hello1"
	}()

	go func() {
		fmt.Println("Writing Hello2 to channel")
		ch <- "Hello2"
	}()

	go func() {
		fmt.Println("Writing Hello3 to channel")
		ch <- "Hello3"
	}()

	go func() {
		fmt.Println("Writing Hello4 to channel")
		ch <- "Hello4"
	}()

	go func() {
		for msg := range ch {
			fmt.Println(msg)
			time.Sleep(2 * time.Millisecond)
		}
	}()

	time.Sleep(10 * time.Millisecond)
}
