package loops

import (
	"fmt"
	"time"
)

func CountDown(sec int) {

	for cur := sec; cur > 0; cur-- {
		fmt.Println(cur)
		time.Sleep(time.Second)
	}

	fmt.Println("Boom !!!")
}
