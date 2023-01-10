package concurrency

import (
	"fmt"
	"time"
)

// func TestConcurrency() {
// 	go One()
// 	go Two()
// }

func One() {
	time.Sleep(3 * time.Second)
	fmt.Println("one")
}

func Two() {
	fmt.Println("two")
}
