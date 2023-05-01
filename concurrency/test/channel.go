package main

import (
	"fmt"
)

func ChannelTest() {

	ch := make(chan int)

	go func() {
		fmt.Println("Writing to channel")
		ch <- 262
	}()

	fmt.Println("Reading from channel")
	msg := <-ch
	fmt.Println(msg)
}
