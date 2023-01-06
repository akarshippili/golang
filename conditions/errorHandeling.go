package conditions

import (
	"fmt"
	"os"
)

func TryRead() {

	_, err := os.Open("./me.txt")

	if err != nil {
		fmt.Println("error: ", err)
	}

}
