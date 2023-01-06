package conditions

import "fmt"

func PrintEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
