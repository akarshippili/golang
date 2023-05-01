package main

import (
	"fmt"
	"time"
)

func SelectChannelTest() {

	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}()

	go func() {
		ch1 <- "Hello"
	}()

	go func() {
		ch2 <- "World"
	}()

	time.Sleep(5 * time.Millisecond)
	fmt.Println("Done!")
}
